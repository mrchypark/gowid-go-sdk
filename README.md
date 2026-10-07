# gowid-go-sdk

Go client for the [Gowid Open API](https://openapi.gowid.com/api-reference).
The official spec snapshot this SDK targets is kept in [spec.md](spec.md).

## Install

```sh
go get github.com/mrchypark/gowid-go-sdk
```

```go
import "github.com/mrchypark/gowid-go-sdk/client"
```

## Usage

The API key is sent as the raw `Authorization` header value, with no `Bearer`
prefix. Request it from Gowid support; it is scoped to a single user.

```go
c := client.NewClient(os.Getenv("API_KEY"))
```

New integrations should use the V2 endpoints. In the official spec 12 of the
14 V1 operations are marked deprecated; `GetMembers` (V1) and `AddComment`
remain current and have no V2 equivalent. `SearchExpensesV2` is itself marked
deprecated; use `GetExpenseStatementsV2` instead.

Dates are `yyyyMMdd`. On `/v2/expense-statements` both bounds are inclusive.

```go
statements, err := c.GetExpenseStatementsV2(&client.ExpenseSearchOptionsV2{
	StartDate: "20260101",
	EndDate:   "20260131",
	Page:      0,
	Size:      20,
})
if err != nil {
	return err
}
for _, s := range statements.Data.Content {
	fmt.Println(s.ExpenseDate, s.KRWAmount, s.StoreName)
}
```

For a single targeted edit, prefer the dedicated setter over the full update:

```go
_, err := c.UpdateExpenseMemoV2(expenseID, "영수증 확인 완료")
_, err = c.UpdateExpensePurposeV2(expenseID, client.UpdatePurposeRequestV2{PurposeID: purposeID})
_, err = c.UpdateExpenseParticipantsV2(expenseID, client.UpdateParticipantsRequestV2{
	ParticipantIDs: []int64{userID},
})
```

## Pagination

`Page` is zero-based. `Size` defaults to 20 server-side and the documented
maximum is 100 on `/v2/expense-statements` and `/v2/cards`; a larger `Size` is
rejected client-side before the request goes out. Iterate `Page` until
`Data.Last` is true.

## Errors

Every response is wrapped in the common envelope
`{"result":{"code":…,"desc":…},"data":…}`, and success is code `20000000`.
The SDK returns an `*client.APIError` when the HTTP status is not 2xx **or** the
result code is not `20000000` — including a failure code returned with HTTP 200,
so a 2xx status alone is not success.

```go
var apiErr *client.APIError
if errors.As(err, &apiErr) {
	log.Println(apiErr.StatusCode, apiErr.Code, apiErr.Desc)
}
```

`Code` and `Desc` are zero/empty when the body was not a Gowid envelope (for
example an HTML error page from a gateway); `Body` then holds the raw payload.
A 2xx response whose envelope is malformed or missing its `result` is a plain
error, not an `APIError`, because there is no result code to report. Typical
codes: `40100010` invalid API key, `40320003` no access to the expense,
`40000001` bad parameter, `40020020` expense not found.

## Updating an expense

`UpdateExpenseV2` is a full update of purpose, memo, participants and
requirement answers in one call. In `client.UpdateExpenseRequestV2` the scalar
fields are pointers: a nil `PurposeID` or `Memo` is omitted from the request
body, while a non-nil pointer to the zero value is sent as `0` or `""`. Nil
slices and maps are omitted; empty (non-nil) slices and maps are sent as empty
arrays and objects.

Those are statements about the wire payload only. The official spec does not
state that omitted fields are retained or that explicit empties clear the
stored value, so treat that as undefined and verify against your own data
before relying on it. When you only want to change one thing, use the
dedicated memo, purpose or participants setter, whose payload is unambiguous.

`expenseId` is always taken from the method's path argument and is not a struct
field, so a mismatched body ID cannot be sent. `isIgnoredFiles` is always sent
as `true`, because the Open API does not support file updates. The nullable
business registration number is `StoreRegistrationNumber *string` on
`client.ExpenseDetailV2`.

## Coverage

All 28 operations in the official spec are implemented: the V2 surface (expense
detail, statements search, deprecated `/v2/expenses` search, not-submitted,
full update, memo, purpose, bulk purpose, participants, approval status, bulk
approve, purposes, purpose requirements, cards) and the V1 surface retained
for existing integrations (members, purposes, purpose requirements, expenses
search, not-submitted, expense detail, memo, purpose, bulk purpose,
participants, full update, approval status, bulk approve, comment).

Two things are unresolved in the official document and are not promised here:
whether `metadata` fields may be null, and the meaning of the two purpose-flag
booleans on the V1 purpose schema. The published schema lists both
`purposeRequirementInputValue` and `isPurposeRequirementInputValue` next to
`purposeRequirementItem` in the update request, and this SDK exposes both as
optional pointers (`PurposeRequirementInputValue`, `IsPurposeRequirementInputValue`).
Whether they are aliases, which one the server reads, or how they interact with
explicit empty values is not stated in the docs. Set neither unless you know
which one your integration expects; there is no default.

## Verifying against the live API

`verify_full.go` is a manual smoke test that runs a selected set of read-only
checks (the list endpoints, which need no IDs of your own) through the SDK and
exits non-zero if any fails. Single-resource reads and writes are not included,
because they would need IDs from your own data. It needs a real key and
network:

```sh
API_KEY=… go run verify_full.go
```

`go test ./...` runs offline against local test servers only.

## Breaking changes

The SDK is in initial v0.x development, so these V1 corrections were made to
match the official spec rather than to preserve an earlier wire contract.
Migrate as follows.

- `GetMembers()` no longer takes an options struct; the spec defines no query
  parameters for it.
- `GetPurposesOptions` no longer has `Limit` or `Page`; `/v1/purposes` accepts
  only `isActivated`.
- Every ID and integral amount the spec declares `int64` is `int64` here
  (member, purpose, category and expense IDs, participant IDs, amounts). They
  were `int`, which cannot represent the full documented range on 32-bit
  platforms; `int64` preserves that range on every platform.
- `GetPurposeRequirements` returns the list at `resp.Data.Content`, matching the
  official `PurposeRequirementResDto` wrapper, instead of a bare slice.
- Single-resource mutations (memo, purpose, participants, full update,
  approval status) now return `ExpenseDetail`. The bulk purpose call returns
  `bool`, and bulk approval returns `ExpenseSimple` items.
- `UpdateExpenseRequest` uses the official field names `ParticipantIdList` and
  `ExternalUserList`. `UpdateParticipantsRequest` uses `ParticipantIds` and
  `ExternalUsers`, matching its own schema.
- Optional scalar request fields are pointers (`*int64`, `*string`, `*bool`) so an
  omitted field and an explicit zero are distinguishable. This applies to the
  approval request too: `ApprovedAmount *int64` and `ApprovedAt *string`.
  Pointer presence only controls what goes on the wire; see the note above
  before relying on it to clear a stored value.
- `Purpose.Requirement` is `*PurposeRequirement`. Per the spec, `purposeId` and
  `options` on the V1 requirement item are write-only and are not returned by
  `GET /v1/purposes`.
- `UpdatePurposeRequest` carries both published purpose-flag booleans
  (`PurposeRequirementInputValue` and `IsPurposeRequirementInputValue`). Neither
  is defaulted or treated as an alias for the other; see the note under Coverage
  for what is still unknown about them.
- The bulk request types are now named slices of the single-item types:
  `UpdateExpensesPurposeRequest` is `[]UpdatePurposeRequest` and
  `ApproveExpensesRequest` is `[]UpdateApprovalStatusRequest`.
- `Do` now returns `*APIError` for a non-success result code even on HTTP 200.
  Code that only checked the HTTP status, or that treated a 2xx as success,
  needs updating.
