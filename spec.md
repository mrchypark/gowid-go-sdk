# API Documentation

## GET {{BASE_URL}}/actuator/health - 서버 헬스 체크

**Update**:

**Description**: 서버의 상태를 확인하는 헬스 체크 API입니다.

### Request

#### Method Type

GET

#### URL

https://openapi.gowid.com/actuator/health

#### Parameters

### Response Example

```JSON
{
    "status": "UP"
}
```

---


## GET {{BASE_URL}}/v1/members - 법인에 소속된 사용자 정보 조회

**Update**: - 응답에서 일부 필드들이 제거 되고, 부서/권한 정보가 새롭게 정의되어 추가 됩니다.

### Request

#### Method Type

GET

#### URL

https://openapi.gowid.com/v1/members

#### Parameters
limit (optional, integer): 현재 서버에서 무시됨(전체 목록 반환)
page (optional, integer): 현재 서버에서 무시됨(전체 목록 반환)
size (optional, integer): 현재 서버에서 무시됨(전체 목록 반환)

### Response Example

```JSON
{
    "result": {
        "code": 20000000,
        "desc": "success"
    },
    "data": [
        {
            "userId": 0,
            "userName": "string",
            "email": "string",
            "isContractor": true,
            "status": "INVITATION",
            "department": { // 부서정보
							"id": 0, // 부서 ID
							"name": "string" // 부서명
						},
            "position": "얏",
            "role": { // 권한 정보
	            "type": "ROLE_MASTER", // role type : ROLE_MASTER, ROLE_MANAGER, ROLE_VIEWER, ROLE_CUSTOM, ROLE_MEMBER
	            "name": "슈퍼관리자", // role 이름
	            "description": null
            },
            "notificationOnOff": true
        }
	]
}
```

---

## GET {{BASE_URL}}/v1/expenses/{expenseId} - 지출내역 단건 조회

**Update**: - [2025-10-10] 이용내역 공제 관련 필드가 추가되었습니다.

### Request

#### Method Type

GET

#### URL

https://openapi.gowid.com/v1/expenses/{expenseId}

#### Parameters

### Response Example

```JSON
{
  "result": {
    "code": 20000000,
    "desc": "success"
  },
  "data": {
    "expenseId": 10145,
    "cardApprovalNumber": "***",
    "expenseDate": "20250821",
    "expenseTime": "232257",
    "card": {
      "cardNumber": "***",
      "cardUser": {
        "userName": "***",
        "email": "***",
        "mobileNumber": "",
        "isInvitedUser": true,
        "isContractor": false,
        "isActivated": true,
        "position": "",
        "activatedAt": "***",
        "deactivatedAt": null,
        "notificationOnOff": true,
        "status": "NORMAL",
        "osType": null
      },
      "alias": "***",
      "limitAmount": 0,
      "usedAmount": 0,
      "remainAmount": 0,
      "companyCode": "0306",
      "namedYn": null,
      "cardName": "고위드 스타트업 신한",
      "cardType": "신한카드",
      "userNm": "***",
      "unmaskedCardNumber": null,
      "duplicationStatus": "NORMAL",
      "fullCardNumber": "***",
      "invalid": false
    },
    "user": {
      "userName": "***",
      "email": "***",
      "mobileNumber": "",
      "isInvitedUser": true,
      "isContractor": false,
      "isActivated": true,
      "position": "",
      "activatedAt": "***",
      "deactivatedAt": null,
      "notificationOnOff": true,
      "status": "NORMAL",
      "osType": null
    },
    "useAmount": 12000.00,
    "currency": "KRW",
    "krwAmount": 12000,
    "approvalStatus": "SUBMITTED",
    "approvedAmount": null,
    "approvedAt": null,
    "approvedBy": null,
    "comments": null,
    "purpose": {
      "name": "야근식비",
      "category": {
        "categoryId": 0,
        "name": "복리후생비"
      },
      "limitAmount": 0,
      "listOrder": 0,
      "isActivated": true,
      "hasRequirement": false,
      "limitType": "PERSON"
    },
    "participants": [
      {
        "userName": "***",
        "email": "***",
        "mobileNumber": "",
        "isInvitedUser": true,
        "isContractor": false,
        "isActivated": true,
        "position": "",
        "activatedAt": "***",
        "deactivatedAt": null,
        "notificationOnOff": true,
        "status": "NORMAL",
        "osType": null
      }
    ],
    "expenseExternalUsers": [],
    "storeName": "***",
    "storeAddress": "***",
    "memo": "10:00",
    "evidenceList": [
      {
        "evidenceId": 12345,
        "fileName": "IMG_20250825142019108_1756099221723.jpeg",
        "mimeType": "image/jpeg",
        "signedUrl": "https://storage.googleapis.com/..."
      }
    ],
    "companyCode": null,
    "commentCount": null,
    "purposeRequirementItem": null,
    "purposeRequirementItemType": null,
    "purposeRequirementValue": null,
    "isPurposeRequirementInputValue": false,
    "expenseDeductionResDto": {
      "isExpenseDeductible": true,  // 이용내역 공제 가능 여부
      "isDeducted": true, // 이용내역 공제 여부
    }
  }
}
```

- evidenceList[].signedUrl 은 제출한 증빙서류 파일의 링크입니다. 본 API 호출로부터 15분간 유효합니다.

---

## GET {{BASE_URL}}/v1/expenses - 지출관리 내역에서 용도, 사용자, 메모에 대한 키워드 검색 및 조회

**Update**: - 요청 파라미터에서 memoName → memo 로 변경 됩니다.

- 요청 파라미터에서 month → startDate 로 변경 됩니다. (월단위 → 시작날짜단위)
- 응답에서 일부 필드가 제거 됩니다.
- 응답에서 용도의 ‘필수항목’ 정보가 추가 됩니다.
- 응답에서 용도의 ‘한도유형’ 정보가 추가 됩니다.

### Request

#### Method Type

GET

#### URL

https://openapi.gowid.com/v1/expenses

#### Parameters
approvalState (optional, string): APPROVED, NOT_SUBMITTED, PARTIAL_APPROVED, REJECTED, SUBMITTED
memo (optional, string): memo, purposeName, userName 중 하나만 사용
purposeName (optional, string): memo, purposeName, userName 중 하나만 사용
userName (optional, string): memo, purposeName, userName 중 하나만 사용
startDate (required, string): yyyyMMdd 또는 yyyy-MM-dd
size (optional, integer): 한 페이지에 조회할 데이터 건수
page (optional, integer): 페이지 인덱스(0부터)

### Response Example

```json
{
  "result": {
    "code": 20000000,
    "desc": "success"
  },
  "data": {
    "totalPages": 70,
    "totalElements": 140,
    "last": false,
    "content": [
      {
        "expenseId": 68516,
        "expenseDate": "20230619",
        "expenseTime": "101949",
        "useAmount": 12160.0,
        "currency": "KRW",
        "krwAmount": 12160,
        "approvedAmount": 12160, // nullable
        "approvedAt": "20230619113332", // nullable
        "approvalStatus": "SUBMITTED",
        "purpose": { // 목적 미지정 시 null
          "purposeId": 2598,
          "name": "야근교통비2",
          "limitType": "ITEM", // 이용내역 용도의 한도 유형
          "limitAmount": 50000,
          "isActivated": true,
          "hasRequirement": true // 이용내역 용도의 필수항목 설정 여부
        },
        "cardAlias": null,
        "cardUserName": "홍길동",
        "shortCardNumber": "7897",
        "storeName": "KT유선상품 자동납부",
        "storeAddress": "서울시 강남구 ...",
        "memo": "",
        "commentCount": 0,
        "evidenceCount": 0,
        "participantCount": 3,
        "representativeParticipant": "dddd",
        "participants": [
          {
            "userId": 9339,
            "userName": "123123123123123123123123123123123123123"
          },
          {
            "userId": 10146,
            "userName": "dddd"
          },
          {
            "userId": 10392,
            "userName": "dd"
          }
        ],
        "expenseExternalUsers": [],
        "purposeRequirementItem": "출근시간", // 이용내역 용도의 필수항목명
        "purposeRequirementItemType": "TEXT", // 이용내역 용도의 필수항목 유형
        "purposeRequirementValue": "8시", // 이용내역 용도의 필수항목값
        "isPurposeRequirementInputValue": false // 이용내역 용도의 필수항목 직접입력 허용 여부
      },
      {
        "expenseId": 68195,
        "expenseDate": "20230614",
        "expenseTime": "121835",
        "useAmount": 11300.0,
        "currency": "KRW",
        "krwAmount": 11300,
        "approvedAmount": 100, // nullable
        "approvedAt": "20230614143346", // nullable
        "approvalStatus": "SUBMITTED",
        "purpose": { // 목적 미지정 시 null
          "purposeId": 2703,
          "name": "용도(수정)",
          "limitType": "PERSON",
          "limitAmount": 0,
          "isActivated": true,
          "hasRequirement": false
        },
        "cardAlias": null,
        "cardUserName": "홍길동",
        "shortCardNumber": "9818",
        "storeName": "카카오페이(택시)",
        "storeAddress": "서울시 ...",
        "memo": "apah",
        "commentCount": 2,
        "evidenceCount": 0,
        "participantCount": 2,
        "representativeParticipant": "123123123123123123123123123123123123123",
        "participants": [
          {
            "userId": 9339,
            "userName": "123123123123123123123123123123123123123"
          },
          {
            "userId": 10187,
            "userName": "123"
          }
        ],
        "expenseExternalUsers": [],
        "purposeRequirementItem": null,
        "purposeRequirementItemType": null,
        "purposeRequirementValue": null,
        "isPurposeRequirementInputValue": false
      }
    ]
  }
}
```

---

## N/A - 미제출 영수증 개수 조회

**Update**: - 미제출 영수증 목록조회 응답에 전체 갯수 필드가 추가 되어 갯수 조회 API 는 제거 됩니다.

### Request

#### Method Type

#### URL

#### Parameters

### Response Example

```JSON

```

---

## GET {{BASE_URL}}/v1/expenses/not-submitted - 미제출 영수증 목록 조회

**Update**: - 용도의 ‘한도유형’ 정보가 응답에 추가 됩니다.

- 응답에서 추천 용도의 필수항목 정보가 추가 됩니다.
- 응답에서 전체 미제출 영수증 갯수, 전체 페이지 수, 마지막 페이지 여부가 추가 됩니다.
- 법인의 전체 미제출 영수증 목록이 조회 됩니다.

### Request

#### Method Type

GET

#### URL

https://openapi.gowid.com/v1/expenses/not-submitted

#### Parameters

size (optional, integer): 한 페이지에 노출할 데이터 건수
page (optional, integer): 조회 페이지 오프셋(0부터)

### Response Example

```JSON
{
    "result": {
        "code": 20000000,
        "desc": "success"
    },
    "data": {
        "content": [
            {
                "expenseId": 0,
                "expenseDate": "string",
                "expenseTime": "string",
                "useAmount": 0.00,
                "currency": "string",
                "krwAmount": 0,
                "storeName": "string",
                "approvalStatus": "NOT_SUBMITTED",
                "cardAlias": "string",
                "shortCardNumber": "string",
                "recommendedPurposeList": [
                    {
                        "purposeId": 0,
                        "name": "string",
                        "category": {
                            "categoryId": 0,
                            "name": "string"
                        },
                        "listOrder": 0,
									      "limitType": "ITEM", // ITEM : 결제건별 한도, PERSONE : 인원별 한도
                        "limitAmount": 0,
                        "isActivated": true,
                        "hasRequirement": true, // true : 필수항목 설정 용도, false : 필수항목 미설정 용도
									      "requirement": { // 용도의 필수항목 정보, hasRequirement 가 false인 경우 "requirement": null
									          "id": 0,
									          "type": "TEXT", // TEXT : 직접입력 방식, SELECT : 목록선택 방식
									          "item": "string", // 필수항목의 제목
									          "guideDesc": "string", // 필수항목이 TEXT type 일 때에 사용되는 가이드 문구
									          "isAvailableInput": false // 필수항목이 SELECT type 일 때에 사용되는 직접입력 허용 여부
									      },
									      "isDeducted": true // 회계 ERP 설정 - true : 해당 용도 공제처리, false : 해당 용도 불공제 처리
                    },
                    {
                        "purposeId": 0,
                        "name": "string",
                        "category": {}, // 계정과목 없을 경우
                        "listOrder": 0,
                        "limitType": "ITEM",
                        "limitAmount": 0,
                        "isActivated": true,
                        "hasRequirement": false,
                        "requirement": {},
                        "isDeducted": false
                    }
                ]
            }
        ],
        "totalPages": 97,
        "totalElements": 1923,
        "last": false
    }
}
```

---


---

## GET {{BASE_URL}}/v1/purposes - 법인에서 정한 사용용도 정책 목록 조회

**Update**: - 용도의 ‘활성여부’ 요청 파라미터가 추가 됩니다.

- 응답에서 용도의 ‘한도유형’ 정보가 추가 됩니다.
- 응답에서 용도의 ‘필수항목’ 정보가 추가 됩니다.

### Request

#### Method Type

GET

#### URL

https://openapi.gowid.com/v1/purposes

#### Parameters

isActivated (optional, boolean): 활성/비활성 용도 필터(미지정 시 전체)
limit (optional, integer): 현재 서버에서 무시됨(전체 목록 반환)

### Response Example

```JSON
{
 "result": {
    "code": 0,
    "desc": "string"
  },
  "data": [
		{
			"purposeId": 0,
      "name": "string",
      "category": {
          "categoryId": 0,
          "name": "string"
      },
      "listOrder": 0,
      "limitType": "ITEM", // ITEM : 결제건별 한도, PERSONE : 인원별 한도
      "limitAmount": 0,
      "isActivated": true,
      "hasRequirement": true, // true : 필수항목 설정 용도, false : 필수항목 미설정 용도
      "requirement": { // 용도의 필수항목 정보, hasRequirement 가 false인 경우 "requirement": null
          "id": 0,
          "type": "TEXT", // TEXT : 직접입력 방식, SELECT : 목록선택 방식
          "item": "string", // 필수항목의 제목
          "guideDesc": "string", // 필수항목이 TEXT type 일 때에 사용되는 가이드 문구
          "isAvailableInput": false // 필수항목이 SELECT type 일 때에 사용되는 직접입력 허용 여부
      },
      "isDeducted": true // 회계 ERP 설정 - true : 해당 용도 공제처리, false : 해당 용도 불공제 처리
  ]
}
```

---

## GET {{BASE_URL}}/v1/purposes/{purposeId}/requirements - 법인에서 정한 사용용도의 선택형 필수항목값 목록조회

**Update**: - 새로 추가된 API 입니다.

- 용도의 선택형 필수항목값 목록조회로, 선택형 필수항목이 설정된 용도의 경우 설정된 항목값들을 조회할 수 있습니다.
- 영수증 제출 시 선택형 필수항목 용도의 필수항목값 선택 시 사용

### Request

#### Method Type

GET

#### URL

https://openapi.gowid.com/v1/purposes/{purposeId}/requirements

#### Parameters

### Response Example

```JSON
{
    "result": {
        "code": 20000000,
        "desc": "success"
    },
    "data": {
        "content": [ // 선택형 필수항목값 목록
            "string",
            "string",
            "string"
        ]
    }
}
```

---

## PUT {{BASE_URL}}/v1/expenses/{expenseId}/purposes - 지출내역 수정 - 용도 정보

**Update**: - RequestBody 에 수정하고자 하는 용도가 필수항목 설정이 되어 있을 경우, 필수항목 정보가 추가 됩니다.

- 응답에서 일부 필드가 제거 됩니다.
- 응답에서 이용내역의 ‘필수항목’ 정보가 추가 됩니다.

### Request

#### Method Type

PUT

#### URL

https://openapi.gowid.com/v1/expenses/{expenseId}/purposes

#### Parameters

##### RequestBody Sample

```JSON
{
    "purposeId" : 2598,
    "purposeRequirementItem": "출근시간",
    "purposeRequirementItemType": "SELECT",
    "purposeRequirementValue": "8시15분",
    "isPurposeRequirementInputValue": true
}
```

### Response Example

```JSON
{
    "result": {
        "code": 0,
        "desc": "string"
    },
    "data": {
        "expenseId": 0,
        "cardApprovalNumber": "string",
        "expenseDate": "string",
        "expenseTime": "string",
        "card": {
            "cardNumber": "string",
            "cardUser": {
                "userId": 0,
                "userName": "string",
                "email": "string",
                "mobileNumber": "string",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "string",
                "imgUrl": "string",
                "activatedAt": "string",
                "deactivatedAt": "string",
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": "string"
            },
            "alias": "string",
            "limitAmount": 0,
            "usedAmount": 0,
            "remainAmount": 0,
            "companyCode": "string",
            "namedYn": "string",
            "cardName": "string",
            "cardType": "string",
            "userNm": "고위드_string",
            "unmaskedCardNumber": "string",
            "duplicationStatus": "NORMAL",
            "fullCardNumber": null,
            "invalid": false
        },
        "user": {
            "userName": "string",
            "email": "string",
            "mobileNumber": "string",
            "isInvitedUser": true,
            "isContractor": false,
            "isActivated": true,
            "position": "string",
            "activatedAt": "string",
            "deactivatedAt": "string",
            "notificationOnOff": true,
            "status": "INVITATION",
            "osType": "string"
        },
        "useAmount": 0,
        "currency": "string",
        "krwAmount": 0,
        "approvalStatus": "SUBMITTED",
        "approvedAmount": 0,
        "approvedAt": "string",
        "approvedBy": "string",
        "purpose": {
            "name": "string2",
            "category": {
                "categoryId": 0,
                "name": "string"
            },
            "limitAmount": 0,
            "listOrder": 0,
            "isActivated": true,
						"hasRequirement": false,
            "limitType": "PERSON"
        },
        "participants": [
            {
                "userName": "name",
                "email": "juns@goaie.com",
                "mobileNumber": "",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "",
                "imgUrl": null,
                "activatedAt": "20221002211939",
                "deactivatedAt": null,
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": null
            }
        ],
        "expenseExternalUsers": [],
        "storeName": "string",
        "storeAddress": "string",
        "memo": "string",
        "expenseEvidences": [],
        "companyCode": "string",
        "commentCount": 0,
        "purposeRequirementItem": "string",
        "purposeRequirementItemType": "TEXT",
        "purposeRequirementValue": "string",
        "isPurposeRequirementInputValue": false
    }
}
```

---

## PUT {{BASE_URL}}/v1/expenses/purposes - 지출내역 수정 - 다건 용도 변경

**Update**: - 용도의 필수항목이 설정되어 있는 경우 RequestBody 에 필수항목 정보가 추가 됩니다.

- 용도의 필수항목이 설정되어 있는 경우 응답에 필수항목 정보가 추가 됩니다.
- 2025-12-29 기준 실제 호출 결과 50000000(서버 에러) 발생.

### Request

#### Method Type

PUT

#### URL

https://openapi.gowid.com/v1/expenses/purposes

#### Parameters

#### Request Body Example

```JSON
[
	{
	    "expenseId": 1,
	    "purposeId": 1,
			"purposeRequirementItem": "출근시간",
	    "purposeRequirementItemType": "SELECT",
	    "purposeRequirementValue": "8시15분",
	    "isPurposeRequirementInputValue": true
	}
]
```

### Response Example

```JSON
{
    "result": {
        "code": 0,
        "desc": "string"
    },
    "data": true
}
```

---

## PUT {{BASE_URL}}/v1/expenses/{expenseId}/memo - 지출내역 수정 - 메모 정보

### Request

#### Method Type

PUT

#### URL

https://openapi.gowid.com/v1/expenses/{expenseId}/memo

#### Parameters

#### Request Body Example

```JSON
{
  "memo": "string"
}
```

### Response Example

```JSON
{
    "result": {
        "code": 0,
        "desc": "string"
    },
    "data": {
        "expenseId": 0,
        "cardApprovalNumber": "string",
        "expenseDate": "string",
        "expenseTime": "string",
        "card": {
            "cardNumber": "string",
            "cardUser": {
                "userId": 0,
                "userName": "string",
                "email": "string",
                "mobileNumber": "string",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "string",
                "imgUrl": "string",
                "activatedAt": "string",
                "deactivatedAt": "string",
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": "string"
            },
            "alias": "string",
            "limitAmount": 0,
            "usedAmount": 0,
            "remainAmount": 0,
            "companyCode": "string",
            "namedYn": "string",
            "cardName": "string",
            "cardType": "string",
            "userNm": "고위드_string",
            "unmaskedCardNumber": "string",
            "duplicationStatus": "NORMAL",
            "fullCardNumber": null,
            "invalid": false
        },
        "user": {
            "userName": "string",
            "email": "string",
            "mobileNumber": "string",
            "isInvitedUser": true,
            "isContractor": false,
            "isActivated": true,
            "position": "string",
            "activatedAt": "string",
            "deactivatedAt": "string",
            "notificationOnOff": true,
            "status": "INVITATION",
            "osType": "string"
        },
        "useAmount": 0,
        "currency": "string",
        "krwAmount": 0,
        "approvalStatus": "SUBMITTED",
        "approvedAmount": 0,
        "approvedAt": "string",
        "approvedBy": "string",
        "purpose": {
            "name": "string2",
            "category": {
                "categoryId": 0,
                "name": "string"
            },
            "limitAmount": 0,
            "listOrder": 0,
            "isActivated": true,
						"hasRequirement": false,
            "limitType": "PERSON"
        },
        "participants": [
            {
                "userName": "name",
                "email": "juns@goaie.com",
                "mobileNumber": "",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "",
                "imgUrl": null,
                "activatedAt": "20221002211939",
                "deactivatedAt": null,
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": null
            }
        ],
        "expenseExternalUsers": [],
        "storeName": "string",
        "storeAddress": "string",
        "memo": "string",
        "expenseEvidences": [],
        "companyCode": "string",
        "commentCount": 0,
        "purposeRequirementItem": "string",
        "purposeRequirementItemType": "TEXT",
        "purposeRequirementValue": "string",
        "isPurposeRequirementInputValue": false
    }
}
```

---

## PUT {{BASE_URL}}/v1/expenses/{expenseId}/participants - 지출내역 수정 - 내/외부 참석자 정보

### Request

#### Method Type

PUT

#### URL

https://openapi.gowid.com/v1/expenses/{expenseId}/participants

#### Parameters

#### Request Body Example

```JSON
{
  "externalUsers": [ // 없을 경우 empty array []
    {
      "company": "string",
      "name": "string"
    }
  ],
  "participantIds": [ // 없을 경우 empty array []
    0
  ]
}
```

### Response Example

```JSON
{
    "result": {
        "code": 0,
        "desc": "string"
    },
    "data": {
        "expenseId": 0,
        "cardApprovalNumber": "string",
        "expenseDate": "string",
        "expenseTime": "string",
        "card": {
            "cardNumber": "string",
            "cardUser": {
                "userId": 0,
                "userName": "string",
                "email": "string",
                "mobileNumber": "string",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "string",
                "imgUrl": "string",
                "activatedAt": "string",
                "deactivatedAt": "string",
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": "string"
            },
            "alias": "string",
            "limitAmount": 0,
            "usedAmount": 0,
            "remainAmount": 0,
            "companyCode": "string",
            "namedYn": "string",
            "cardName": "string",
            "cardType": "string",
            "userNm": "고위드_string",
            "unmaskedCardNumber": "string",
            "duplicationStatus": "NORMAL",
            "fullCardNumber": null,
            "invalid": false
        },
        "user": {
            "userName": "string",
            "email": "string",
            "mobileNumber": "string",
            "isInvitedUser": true,
            "isContractor": false,
            "isActivated": true,
            "position": "string",
            "activatedAt": "string",
            "deactivatedAt": "string",
            "notificationOnOff": true,
            "status": "INVITATION",
            "osType": "string"
        },
        "useAmount": 0,
        "currency": "string",
        "krwAmount": 0,
        "approvalStatus": "SUBMITTED",
        "approvedAmount": 0,
        "approvedAt": "string",
        "approvedBy": "string",
        "purpose": {
            "name": "string2",
            "category": {
                "categoryId": 0,
                "name": "string"
            },
            "limitAmount": 0,
            "listOrder": 0,
            "isActivated": true,
						"hasRequirement": false,
            "limitType": "PERSON"
        },
        "participants": [
            {
                "userName": "name",
                "email": "juns@goaie.com",
                "mobileNumber": "",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "",
                "imgUrl": null,
                "activatedAt": "20221002211939",
                "deactivatedAt": null,
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": null
            }
        ],
        "expenseExternalUsers": [],
        "storeName": "string",
        "storeAddress": "string",
        "memo": "string",
        "expenseEvidences": [],
        "companyCode": "string",
        "commentCount": 0,
        "purposeRequirementItem": "string",
        "purposeRequirementItemType": "TEXT",
        "purposeRequirementValue": "string",
        "isPurposeRequirementInputValue": false
    }
}
```

---

## PUT {{BASE_URL}}/v1/expenses/{expenseId} - 지출내역 수정 - 일반 정보

**Update**: - 용도의 필수항목이 설정되어 있는 경우 RequestBody 에 필수항목 정보가 추가 됩니다.

- 용도의 필수항목이 설정되어 있는 경우 응답에 필수항목 정보가 추가 됩니다.

### Request

#### Method Type

PUT

#### URL

https://openapi.gowid.com/v1/expenses/{expenseId}

#### Parameters

‣

‣

‣

3가지 API 의 통합 버전 API 입니다.

Request 및 Response 에 대해서 참고 부탁 드립니다.

#### Request Body Example

```JSON
{
  "externalUserList": [// 없을 경우 empty array []
    {
      "company": "string",
      "name": "string"
    }
  ],
  "memo": "string",
  "participants": [// 없을 경우 empty array []
    0
  ],
  "purposeId": 0,
	"purposeRequirementItem": "string",
  "purposeRequirementItemType": "string",
  "purposeRequirementValue": "string",
  "isPurposeRequirementInputValue": true
}
```

### Response Example

```JSON
{
    "result": {
        "code": 0,
        "desc": "string"
    },
    "data": {
        "expenseId": 0,
        "cardApprovalNumber": "string",
        "expenseDate": "string",
        "expenseTime": "string",
        "card": {
            "cardNumber": "string",
            "cardUser": {
                "userId": 0,
                "userName": "string",
                "email": "string",
                "mobileNumber": "string",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "string",
                "imgUrl": "string",
                "activatedAt": "string",
                "deactivatedAt": "string",
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": "string"
            },
            "alias": "string",
            "limitAmount": 0,
            "usedAmount": 0,
            "remainAmount": 0,
            "companyCode": "string",
            "namedYn": "string",
            "cardName": "string",
            "cardType": "string",
            "userNm": "고위드_string",
            "unmaskedCardNumber": "string",
            "duplicationStatus": "NORMAL",
            "fullCardNumber": null,
            "invalid": false
        },
        "user": {
            "userName": "string",
            "email": "string",
            "mobileNumber": "string",
            "isInvitedUser": true,
            "isContractor": false,
            "isActivated": true,
            "position": "string",
            "activatedAt": "string",
            "deactivatedAt": "string",
            "notificationOnOff": true,
            "status": "INVITATION",
            "osType": "string"
        },
        "useAmount": 0,
        "currency": "string",
        "krwAmount": 0,
        "approvalStatus": "SUBMITTED",
        "approvedAmount": 0,
        "approvedAt": "string",
        "approvedBy": "string",
        "purpose": {
            "name": "string2",
            "category": {
                "categoryId": 0,
                "name": "string"
            },
            "limitAmount": 0,
            "listOrder": 0,
            "isActivated": true,
						"hasRequirement": false,
            "limitType": "PERSON"
        },
        "participants": [
            {
                "userName": "name",
                "email": "juns@goaie.com",
                "mobileNumber": "",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "",
                "imgUrl": null,
                "activatedAt": "20221002211939",
                "deactivatedAt": null,
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": null
            }
        ],
        "expenseExternalUsers": [],
        "storeName": "string",
        "storeAddress": "string",
        "memo": "string",
        "expenseEvidences": [],
        "companyCode": "string",
        "commentCount": 0,
        "purposeRequirementItem": "string",
        "purposeRequirementItemType": "TEXT",
        "purposeRequirementValue": "string",
        "isPurposeRequirementInputValue": false
    }
}
```

---

## PUT {{BASE_URL}}/v1/expenses/{expenseId}/approval-status - 지출내역 수정 - 승인 상태 정보

**Update**: - RequestBody 에서 approvedAt 필드가 추가 됩니다.
- 승인 처리 시 응답의 approvedAt 값이 서버 시간으로 갱신될 수 있습니다.

### Request

#### Method Type

PUT

#### URL

https://openapi.gowid.com/v1/expenses/{expenseId}/approval-status

#### Parameters

#### Request Body Example

```JSON
{
  "approvalStatus": "APPROVED",
  "approvedAmount": 0,
  "approvedAt": "string"
}
```

### Response Example

```JSON
{
    "result": {
        "code": 0,
        "desc": "string"
    },
    "data": {
        "expenseId": 0,
        "cardApprovalNumber": "string",
        "expenseDate": "string",
        "expenseTime": "string",
        "card": {
            "cardNumber": "string",
            "cardUser": {
                "userId": 0,
                "userName": "string",
                "email": "string",
                "mobileNumber": "string",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "string",
                "imgUrl": "string",
                "activatedAt": "string",
                "deactivatedAt": "string",
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": "string"
            },
            "alias": "string",
            "limitAmount": 0,
            "usedAmount": 0,
            "remainAmount": 0,
            "companyCode": "string",
            "namedYn": "string",
            "cardName": "string",
            "cardType": "string",
            "userNm": "고위드_string",
            "unmaskedCardNumber": "string",
            "duplicationStatus": "NORMAL",
            "fullCardNumber": null,
            "invalid": false
        },
        "user": {
            "userName": "string",
            "email": "string",
            "mobileNumber": "string",
            "isInvitedUser": true,
            "isContractor": false,
            "isActivated": true,
            "position": "string",
            "activatedAt": "string",
            "deactivatedAt": "string",
            "notificationOnOff": true,
            "status": "INVITATION",
            "osType": "string"
        },
        "useAmount": 0,
        "currency": "string",
        "krwAmount": 0,
        "approvalStatus": "SUBMITTED",
        "approvedAmount": 0,
        "approvedAt": "string",
        "approvedBy": "string",
        "purpose": {
            "name": "string2",
            "category": {
                "categoryId": 0,
                "name": "string"
            },
            "limitAmount": 0,
            "listOrder": 0,
            "isActivated": true,
						"hasRequirement": false,
            "limitType": "PERSON"
        },
        "participants": [
            {
                "userName": "name",
                "email": "juns@goaie.com",
                "mobileNumber": "",
                "isInvitedUser": true,
                "isContractor": false,
                "isActivated": true,
                "position": "",
                "imgUrl": null,
                "activatedAt": "20221002211939",
                "deactivatedAt": null,
                "notificationOnOff": true,
                "status": "INVITATION",
                "osType": null
            }
        ],
        "expenseExternalUsers": [],
        "storeName": "string",
        "storeAddress": "string",
        "memo": "string",
        "expenseEvidences": [],
        "companyCode": "string",
        "commentCount": 0,
        "purposeRequirementItem": "string",
        "purposeRequirementItemType": "TEXT",
        "purposeRequirementValue": "string",
        "isPurposeRequirementInputValue": false
    }
}
```

---

## PATCH {{BASE_URL}}/v1/expenses/approval-status/approved - 지출내역 수정 - 다건 승인

**Update**: - RequestBody 에서 단건 승인 상태 변경의 형태를 array 로 전달 주시면 됩니다.

- 응답에서 성공한 이용내역과 실패한 이용내역을 전달 드립니다.

### Request

#### Method Type

PATCH

#### URL

https://openapi.gowid.com/v1/expenses/approval-status/approved

#### Parameters

‣ 요청의 array 형태

#### Request Body Example

```JSON
[
	{
	  "expenseId": 1234567,
	  "approvalStatus": "APPROVED",
	  "approvedAmount": 0,
	  "approvedAt": "string"
	}
]
```

### Response Example

```JSON
{
    "result": {
        "code": 0,
        "desc": "string"
    },
    "data": {
        "succeedStatements": [
            {
                "expenseId": 0,
                "expenseDate": "string",
                "expenseTime": "string",
                "useAmount": 0
                "currency": "string",
                "krwAmount": 0,
                "storeName": "string",
                "approvalStatus": "string",
                "cardAlias": "string",
                "shortCardNumber": "string"
            }
        ],
        "failedStatements": []
    }
}
```

---

## POST {{BASE_URL}}/v1/expenses/{expenseId}/comments - 지출내역 댓글 추가

**Update**: - RequestBody 내 expenseId 필드가 제거 되었습니다.

### Request

#### Method Type

POST

#### URL

https://openapi.gowid.com/v1/expenses/{expenseId}/comments

#### Parameters

#### Request Body Example

```JSON
{
  "comment": "string"
}
```

### Response Example

```JSON
{
    "result": {
        "code": 20000000,
        "desc": "success"
    },
    "data": {
        "commentId": null,
        "author": "string",
        "content": "string",
        "createdAt": "2025-12-29T18:59:54.679"
    }
}
```

---

## Verification Log (2025-12-25)

### /actuator/health

- **Status**: Verified. Returns `{"status":"UP"}`.

### /v1/members

- **Status**: Verified (200 OK).
- **Discrepancy**: Spec defines `role` fields as `roleType`, `roleName`. Actual response uses `type`, `name`. Code matches actual response.

### /v1/purposes

- **Status**: Verified (200 OK). Matches Spec.
