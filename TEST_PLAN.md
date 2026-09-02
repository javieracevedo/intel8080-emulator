# Test Plan — 8-bit Arithmetic Instructions

Scope: the instructions in `cpu/al8.go` — `ADD_X`, `ADD_M_X`, `ADC_X` — plus the flag
machinery they depend on (`SetAddConditionFlags`, `parity8`) and their dispatch through
`Execute`. This is a checklist of what still needs covering, not an implementation.

## Where coverage stands today

- `TestADD_X` runs all seven registers but initializes every one of them to `0x1`, so no
  case distinguishes one register from another and no case sets any flag. It never asserts
  flags at all.
- `TestADD_M_X` is a single case: one HL value, one memory value, no flag assertions.
- `TestADC_X` covers all seven registers across three groups (no carry-in, carry-in with
  overflow, carry-in without overflow) and asserts the full flag byte via `assertFlags`.
  This is the shape the other two should grow into.
- `SetAddConditionFlags`, `parity8`, and the `Execute` dispatch for opcodes `0x80`–`0x8F`
  have no direct tests.

## Shared flag matrix

Every arithmetic instruction ends in `SetAddConditionFlags`, so the same set of flag
outcomes needs exercising once per instruction rather than being assumed from `ADC_X`.
Each of these should assert the complete `Flags` byte, not just the flag under test.

- No flags set: `0x01 + 0x00` → `0x01`. Nothing set, including parity (odd).
- Parity alone: `0x01 + 0x02` → `0x03`. `P` only.
- Sign alone: `0x80 + 0x03` → `0x83`. `S` only; result has odd parity and no carries.
- Aux carry alone: `0x0F + 0x01` → `0x10`. `AC` only; low nibble crosses, byte does not.
- Carry alone: `0xF0 + 0x23` → `0x13`. `CY` only; byte carries out, low nibble does not.
- Zero: `0xFF + 0x01` → `0x00`. `CY`, `AC`, `Z` and `P` together. Worth an explicit note in
  the test that `Z` can never appear alone — a zero result always has even parity, so `P`
  rides along with it.
- Maximum simultaneous flags: `0xFF + 0xBF` → `0xBE`, giving `CY | S | AC | P`. `S` and `Z`
  are mutually exclusive, so this is the widest combination reachable.
- Carry boundary: `0xFE + 0x01` → `0xFF` (`S | P`, no `CY`), against `0xFF + 0x01` → `0x00`
  with `CY` set. Pins the comparison at `> 0xFF` rather than `>=`.
- Aux carry boundary: `0x0E + 0x01` → `0x0F` (`P` only, no `AC`), against `0x0F + 0x01` →
  `0x10` with `AC` set. Pins the nibble comparison at `> 0x0F`.
- Sign boundary: `0x7F + 0x00` → `0x7F` with no flags, against `0x7F + 0x01` → `0x80` with
  `S | AC`.
- Flags are cleared, not just set: pre-set all five flags on the CPU, run an operation that
  should produce none, and assert `Flags` comes back `0x00`. `Init` does not reset `Flags`,
  so this catches a missing `ClearFlag` branch.
- No stray bits: after any operation, assert `Flags &^ (CY|Z|S|P|AC) == 0`. Bits `0x40`,
  `0x10` and `0x04` are unused in this layout and nothing should ever write them.

## ADD_X

- Per-register coverage with **distinct** values in every register, so a case that reads the
  wrong register produces a different result. The current all-`0x1` fixture cannot fail that
  way.
- The flag matrix above, driven through the operand register rather than through `A`.
- A case proving flags come from the operand register, not a hardcoded one: `A = 0x01`,
  `B = 0x00`, `C = 0x0F`, then `ADD_X(C)` → `A = 0x10` with `AC` set. This currently fails —
  `cpu/al8.go:10` passes `c.REGISTERS[B]` to `SetAddConditionFlags` instead of
  `c.REGISTERS[x]`, the same hardcoded-`B` bug that was fixed in `ADC_X` in `f6bf88b`.
- `ADD A` doubling: `A = 0x9A` → `0x34` with `CY` and `AC` set. The accumulator is both
  operands, so a test that reads a stale `A` for one side shows up here.
- Carry-in is ignored: pre-set `CY`, then `A = 0x01`, `B = 0x01`, `ADD_X(B)` → `A = 0x02`
  (not `0x03`) and `CY` ends clear. This is the whole difference between `ADD` and `ADC`
  and nothing asserts it right now.
- Wraparound preserves only the low byte: `0xFF + 0xFF` → `0xFE` with `CY | S | AC`.

## ADD_M_X

- Address assembly from the H/L pair. At minimum `H = 0x12, L = 0x34` → address `0x1234`,
  so a swapped MSB/LSB fails. The current single case uses `H = L = 0xFF`, which is
  symmetric and cannot detect a swap.
- Address boundaries: `HL = 0x0000`, `HL = 0xFFFF`, and one mid-range value.
- The full flag matrix, with the second operand coming from memory instead of a register.
- Reading a zero byte from untouched memory: `A = 0x05`, `HL` pointing at an address never
  written → `A` unchanged at `0x05`, and flags recomputed accordingly.
- Accumulator semantics. The 8080's `ADD M` is `A ← A + (HL)`, but `ADD_M_X` sums
  `REGISTERS[x] + memory[addr]` and stores into `A`, so for any `x` other than `A` it is not
  `ADD M`. `Execute` only ever calls `ADD_M_X(A)` (opcode `0x86`), where the two agree. The
  existing test calls `ADD_M_X(B)` and locks in the non-8080 behavior. Decide whether the
  parameter should exist at all; whichever way it goes, a test should state the intent
  rather than leave both readings valid.
- Memory is a package-level global (`memory.MEMORY`), shared by every test in the package.
  Each test needs `t.Cleanup` to restore the bytes it wrote, and cases should use distinct
  addresses so they cannot collide. Nothing here is safe for `t.Parallel` until that holds.

## ADC_X

Already the best-covered of the three. Remaining gaps:

- Carry-in as the sole cause of a flag crossing, which `ADD` can never reach:
  - `A = 0x0F`, operand `0x00`, `CY` set → `0x10` with `AC` set.
  - `A = 0xFF`, operand `0x00`, `CY` set → `0x00` with `CY`, `AC`, `Z`, `P` set.
  - `A = 0x7F`, operand `0x00`, `CY` set → `0x80` with `S` and `AC` set.
- Carry is read before it is recomputed: `A = 0xFF`, operand `0x01`, `CY` set → `0x01`, and
  `CY` ends set from *this* addition. A test that only checks the final flag can't tell a
  correct read-then-write from a double count; assert the result value too.
- `ADC A` with carry-in across the overflow boundary: `A = 0x80`, `CY` set → `0x01` with
  `CY` set.
- The `S`-alone and `CY`-alone rows from the flag matrix, which the current three groups
  don't reach.

## SetAddConditionFlags

Worth testing directly rather than only through the instructions, since all three share it.

- The `carry` parameter contributes to the sum, the aux-carry nibble sum, and nothing else.
- Called twice in a row with different inputs, the second call fully determines `Flags` —
  no residue from the first.
- The whole flag matrix above, called directly with explicit `(a, b, carry)` triples. This
  is the cheapest place to get exhaustive coverage; the per-instruction tests then only need
  enough cases to prove they pass the right operands through.

## parity8

- `0x00` → true (zero bits set counts as even).
- `0xFF` → true, eight bits.
- `0x01` → false, `0x80` → false — single bit at each end.
- `0x03` → true, `0x7F` → false (seven bits).
- Confirm the 8080 convention this encodes: parity flag *set* means even parity.

## Execute dispatch and cycle counts

None of this is covered, and there are real gaps behind it.

- `0x80`–`0x85` and `0x87` each reach `ADD_X` with the correct register. A table over
  opcodes with distinct register values catches a mis-mapped case.
- `0x86` reaches `ADD_M_X` and reads through HL.
- `0x88`–`0x8F` (`ADC B` through `ADC A`, and `ADC M`) are **absent from the switch** in
  `cpu/instructions.go` — `ADC_X` is implemented and tested but unreachable through
  `Execute`. Tests should cover them once the cases are added.
- `ADC M` (`0x8E`) has no implementation at all. It needs an `ADC_M_X` equivalent plus the
  same address-assembly and flag coverage as `ADD_M_X`.
- `CYCLES_TABLE` has no entries above `0x7F`, so every `ADD`/`ADC` currently adds zero
  cycles. Tests should assert `CyclesCount` advances by 4 for the register forms and 7 for
  the memory forms.
- An unknown opcode falls through the switch silently while still adding its table cycles.
  Decide what that should do and assert it.

## Test hygiene worth encoding

- `Init` does not clear `Flags`. Every table row must zero it explicitly, the way `TestADC_X`
  does, or state leaks between subtests in declaration order.
- Prefer `Errorf` over `Fatalf` in flag assertions so a single bad case reports every
  mismatched flag plus the wrong result in one run.
- Keep fixture registers distinct per case. Uniform fixtures are the reason the `ADD_X` bug
  above is still passing its test.
