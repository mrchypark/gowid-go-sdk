# GOWID 공식 OpenAPI 명세

- 공식 API 참고 문서: https://openapi.gowid.com/api-reference
- OpenAPI 원본: https://openapi.gowid.com/api-docs
- 스냅샷 조회 시각: 2026-10-07 21:36:48 KST (UTC+09:00) — Asia/Seoul

이 문서는 위 조회 시각에 받은 공식 명세의 스냅샷이며, 실시간으로 갱신되지 않습니다. 원본 JSON의 값은 모두 보존하고 jq로 들여쓰기만 정리했습니다.

```json
{
  "openapi": "3.0.1",
  "info": {
    "title": "고위드 오픈 API (Gowid Open API)",
    "description": "고위드(Gowid) 법인 지출관리 데이터를 외부에서 조회·관리할 수 있는 파트너용 REST API입니다.\n지출내역(이용내역) 조회, 용도 설정, 승인 처리, 멤버 조회 등 지출관리 핵심 기능을 제공합니다.\n\n## 인증 (Authentication)\n\n모든 요청은 `Authorization` 헤더에 고위드로부터 발급받은 **API Key** 를 담아야 합니다.\n`Bearer` 등 접두사는 붙이지 않고 키 값만 그대로 전달합니다.\n\n```\nAuthorization: {발급받은 API Key}\n```\n\n- API Key는 **사용자 단위**로 발급되며, 키에 매핑된 사용자의 권한 범위 내에서만 데이터에 접근할 수 있습니다.\n- 유효하지 않거나 비활성화된 키로 요청하면 `40100010` (유효하지 않은 API Key) 오류가 반환됩니다.\n- API Key 발급·재발급이 필요하면 고위드 담당자에게 문의해 주세요.\n\n## Base URL\n\n```\nhttps://openapi.gowid.com\n```\n\n## 공통 응답 형식\n\n모든 응답은 아래와 같은 공통 envelope으로 감싸여 반환됩니다.\n\n```json\n{\n  \"result\": { \"code\": 20000000, \"desc\": \"success\" },\n  \"data\": { /* 실제 응답 본문 */ }\n}\n```\n\n- `result.code` : 처리 결과 코드. 성공은 `20000000` 입니다.\n- `result.desc` : 결과 메시지.\n- `data` : 엔드포인트별 실제 응답 본문.\n\n## 날짜·시간 포맷\n\n별도 표기가 없으면 문자열(String) 기반의 아래 포맷을 사용합니다.\n\n| 종류 | 포맷 | 예시 |\n|------|------|------|\n| 일자 | `yyyyMMdd` | `20260108` |\n| 시각 | `HHmmss` | `103024` |\n| 일시 | `yyyyMMddHHmmss` | `20260108150121` |\n\n## API 버전 정책 (V1 / V2)\n\n현재 V1과 V2가 함께 제공됩니다. **신규 연동은 V2 사용을 권장합니다.**\nV1에서 V2로 대체된 엔드포인트는 `deprecated`로 표시되어 있으며(취소선), 향후 지원이 종료될 수 있습니다.\n\nV2의 주요 개선점:\n- 기간 검색 지원(`startDate` + `endDate`)\n- 용도 세부입력항목의 복수 선택(SELECT_MULTI) 및 다중 항목 응답 지원\n\n## 주요 오류 코드\n\n오류 시 HTTP 상태 코드와 함께 `result.code`로 세부 사유를 구분합니다.\n코드 형식은 `HTTP상태(3자리) + 서비스(1) + 분류(2) + 세부(2)` 입니다.\n\n| code | 의미 |\n|------|------|\n| `40000001` | 잘못된 파라미터 |\n| `40020018` | 승인 금액이 잘못됨 |\n| `40020019` | 이미 승인 상태가 변경된 이용내역 |\n| `40020020` | 해당 ID로 이용내역을 찾을 수 없음 |\n| `40020021` | 변경할 수 없는 이용내역 상태 |\n| `40020022` | 본인 이용내역은 처리할 수 없음 |\n| `40100010` | 유효하지 않은 API Key |\n| `40320003` | 이용내역에 접근할 권한이 없음 |\n| `50000000` | 서버 오류 |\n\n## 문의\n\n연동 중 문의 사항은 openapi@gowid.com 으로 연락해 주세요.",
    "contact": {
      "name": "고위드 오픈 API 지원",
      "email": "openapi@gowid.com"
    },
    "version": "1.0"
  },
  "servers": [
    {
      "url": "https://openapi.gowid.com",
      "description": "운영(Production)"
    }
  ],
  "security": [
    {
      "ApiKeyAuth": []
    }
  ],
  "tags": [
    {
      "name": "지출내역 V2",
      "description": "지출내역(이용내역)의 조회·수정·승인 등 핵심 기능 (V2, 권장)"
    },
    {
      "name": "용도 V2",
      "description": "법인에 설정된 사용용도 및 세부입력항목 조회 (V2, 권장)"
    },
    {
      "name": "멤버",
      "description": "법인에 소속된 사용자(멤버) 조회"
    },
    {
      "name": "댓글 관리",
      "description": "지출내역에 대한 댓글 등록"
    },
    {
      "name": "지출내역",
      "description": "지출내역(이용내역) 관련 기능 (V1, deprecated — V2 사용 권장)"
    },
    {
      "name": "용도",
      "description": "사용용도 조회 (V1, deprecated — V2 사용 권장)"
    }
  ],
  "paths": {
    "/v2/expenses/{expenseId}": {
      "get": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 단건 조회",
        "description": "특정 지출내역의 상세 정보를 조회합니다.",
        "operationId": "getExpense",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDtoV2"
                }
              }
            }
          }
        }
      },
      "put": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 정보 수정",
        "description": "특정 지출내역의 용도·메모·참석자 등 정보를 한 번에 수정합니다.\n\n- 요청 본문의 `expenseId`는 경로 변수와 일치해야 합니다.",
        "operationId": "updateExpense",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpenseUpdateReqDtoV2"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/expenses/{expenseId}/purposes": {
      "put": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 용도 수정",
        "description": "특정 지출내역의 용도와 용도 세부입력항목 답변을 수정합니다.\n\n- `purposeRequirementAnswerMap`은 `세부입력항목 ID → 답변값 목록(List)` 형태이며, SELECT_MULTI(복수 선택)도 지원합니다.\n- 미제출 내역에 용도를 지정하면 제출 상태가 됩니다.",
        "operationId": "updateExpensePurpose",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpensePurposeUpdateRequestDtoV2"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/expenses/{expenseId}/participants": {
      "put": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 참석자 수정",
        "description": "특정 지출내역의 참석자 정보를 수정합니다. 내부 참석자(`participantIds`)와 외부 참석자(`externalUsers`)를 함께 갱신합니다.",
        "operationId": "updateExpenseParticipants",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpenseParticipantsUpdateRequestDtoV2"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/expenses/{expenseId}/memo": {
      "put": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 메모 수정",
        "description": "특정 지출내역의 메모를 수정합니다.",
        "operationId": "updateExpenseMemo",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpenseMemoReqDtoV2"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/expenses/{expenseId}/approval-status": {
      "put": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 승인상태 변경",
        "description": "특정 지출내역의 승인상태를 변경합니다. 요청 가능한 `approvalStatus`는 `APPROVED`(승인)와 `REJECTED`(반려)입니다.\n\n- **부분승인은 직접 지정하지 않습니다.** `APPROVED`로 요청하면서 전체 금액보다 작은 `approvedAmount`를 보내면 서버가 자동으로 부분승인(`PARTIAL_APPROVED`)으로 처리합니다.\n- `approvedAmount`를 생략하거나 전체 금액 이상으로 보내면 전액 승인됩니다.\n- `NOT_SUBMITTED`/`SUBMITTED` 등 그 외 상태로의 변경 요청은 오류가 발생합니다.\n- 본인 이용내역 자가 승인은 권한에 따라 제한될 수 있습니다.",
        "operationId": "updateExpenseApprovalStatus",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpenseApproveReqDtoV2"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/expenses/purposes": {
      "put": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 용도 일괄 수정",
        "description": "여러 지출내역에 **동일한 용도**를 한 번에 적용합니다.\n\n- `expenseIds`(필수)에 적용할 지출내역 ID 목록을 전달하며, 중복 ID가 있으면 오류가 발생합니다.\n- 모든 대상에 공통 `purposeId`와 `purposeRequirementAnswerMap`이 적용됩니다.",
        "operationId": "updateExpensesPurpose",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpensePurposeBulkUpdateRequestDtoV2"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseBoolean"
                }
              }
            }
          }
        }
      }
    },
    "/v1/expenses/{expenseId}": {
      "get": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 단건 조회",
        "description": "[Deprecated] V2(`GET /v2/expenses/{expenseId}`) 사용을 권장합니다.\n\n특정 지출내역의 상세 정보를 조회합니다.",
        "operationId": "getExpense_1",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDto"
                }
              }
            }
          }
        },
        "deprecated": true
      },
      "put": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 정보 수정",
        "description": "[Deprecated] V2(`PUT /v2/expenses/{expenseId}`) 사용을 권장합니다.\n\n특정 지출내역의 용도·메모·참석자 등 정보를 수정합니다. 요청 본문의 `expenseId`는 경로 변수와 일치해야 합니다.",
        "operationId": "updateExpense_1",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpenseUpdateReqDto"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/expenses/{expenseId}/purposes": {
      "put": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 용도 수정",
        "description": "[Deprecated] V2(`PUT /v2/expenses/{expenseId}/purposes`) 사용을 권장합니다. V2는 세부입력항목 복수 선택(SELECT_MULTI)을 지원합니다.\n\n특정 지출내역의 용도를 수정합니다.",
        "operationId": "updateExpensePurpose_1",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpensePurposeUpdateRequestDto"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/expenses/{expenseId}/participants": {
      "put": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 참석자 수정",
        "description": "[Deprecated] V2(`PUT /v2/expenses/{expenseId}/participants`) 사용을 권장합니다.\n\n특정 지출내역의 내부/외부 참석자 정보를 수정합니다.",
        "operationId": "updateExpenseParticipants_1",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpenseParticipantsUpdateRequestDto"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/expenses/{expenseId}/memo": {
      "put": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 메모 수정",
        "description": "[Deprecated] V2(`PUT /v2/expenses/{expenseId}/memo`) 사용을 권장합니다.\n\n특정 지출내역의 메모를 수정합니다.",
        "operationId": "updateExpenseMemo_1",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpenseMemoReqDto"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/expenses/{expenseId}/approval-status": {
      "put": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 승인상태 변경",
        "description": "[Deprecated] V2(`PUT /v2/expenses/{expenseId}/approval-status`) 사용을 권장합니다.\n\n특정 지출내역의 승인상태를 변경합니다. `APPROVED`(승인)/`REJECTED`(반려)를 보낼 수 있으며, `APPROVED` + 전체 금액보다 작은 `approvedAmount` 조합 시 부분승인(`PARTIAL_APPROVED`)으로 자동 처리됩니다.",
        "operationId": "updateExpenseApprovalStatus_1",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/ExpenseApproveReqDto"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseDetailResDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/expenses/purposes": {
      "put": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 용도 일괄 수정",
        "description": "[Deprecated] V2(`PUT /v2/expenses/purposes`) 사용을 권장합니다.\n\n여러 지출내역의 용도를 일괄 수정합니다. (V1은 항목별로 서로 다른 용도를 지정하는 방식)",
        "operationId": "updateExpensesPurpose_1",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/ExpensePurposeUpdateRequestDto"
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseBoolean"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/expenses/{expenseId}/comments": {
      "post": {
        "tags": [
          "댓글 관리"
        ],
        "summary": "지출내역 댓글 등록",
        "description": "특정 지출내역에 댓글을 등록합니다. 등록 시 댓글 수신 대상자에게 알림이 발송될 수 있습니다.",
        "operationId": "addComment",
        "parameters": [
          {
            "name": "expenseId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "$ref": "#/components/schemas/CommentReqDto"
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseCommentResDto"
                }
              }
            }
          }
        }
      }
    },
    "/v2/expenses/approval-status/approved": {
      "patch": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 일괄 승인·반려",
        "description": "여러 지출내역의 승인상태를 한 번에 변경합니다. 각 항목의 부분승인 규칙은 단건 승인 API와 동일합니다.",
        "operationId": "approveExpenses",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/ExpenseApproveReqDtoV2"
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseStatementBulkApproveResDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v1/expenses/approval-status/approved": {
      "patch": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 일괄 승인·반려",
        "description": "[Deprecated] V2(`PATCH /v2/expenses/approval-status/approved`) 사용을 권장합니다.\n\n여러 지출내역의 승인상태를 한 번에 변경합니다.",
        "operationId": "approveExpenses_1",
        "requestBody": {
          "content": {
            "application/json": {
              "schema": {
                "type": "array",
                "items": {
                  "$ref": "#/components/schemas/ExpenseApproveReqDto"
                }
              }
            }
          },
          "required": true
        },
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseStatementBulkApproveResDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v2/purposes": {
      "get": {
        "tags": [
          "용도 V2"
        ],
        "summary": "사용용도 목록 조회",
        "description": "법인에 설정된 사용용도(purpose) 목록을 조회합니다.\n\n- `isActivated`: `true`이면 활성 용도만, `false`이면 비활성 용도만, 생략 시 전체를 조회합니다.\n- 각 용도는 한도(`limitAmount`/`limitType`), 공제 여부(`isDeducted`), 세부입력항목(`requirements`) 정보를 포함합니다.",
        "operationId": "getPurposes",
        "parameters": [
          {
            "name": "isActivated",
            "in": "query",
            "required": false,
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseListPurposeDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/purposes/{purposeId}/requirements/{requirementId}": {
      "get": {
        "tags": [
          "용도 V2"
        ],
        "summary": "용도 세부입력항목 선택지 목록 조회",
        "description": "특정 용도(`purposeId`)의 특정 세부입력항목(`requirementId`)에 대한 선택지 값 목록을 조회합니다.\n\n- **SELECT / SELECT_MULTI(선택형) 타입 항목에서만** 선택지가 반환됩니다. TEXT(자유입력) 항목은 빈 목록입니다.\n- 비활성 용도이거나 해당 항목이 선택형이 아니면 빈 목록을 반환합니다.",
        "operationId": "getPurposeRequirements",
        "parameters": [
          {
            "name": "purposeId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          },
          {
            "name": "requirementId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponsePurposeRequirementResDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/expenses": {
      "get": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 검색·목록 조회",
        "description": "[Deprecated] V2(`GET /v2/expense-statements`) 사용을 권장합니다.\n\n용도·사용자·메모 키워드로 지출내역을 검색합니다. 키워드가 없으면 해당 기간 내 전체 목록을 반환합니다.\n\n- 종료일(`endDate`)을 입력받지 않으며, 조회 기간은 `startDate` 기준 약 1개월로 **자동 보정**됩니다. 종료일 당일은 **제외**(exclusive)됩니다.",
        "operationId": "searchExpenseHistory",
        "parameters": [
          {
            "name": "criteria",
            "in": "query",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/ExpenseSearchCriteriaV2"
            }
          },
          {
            "name": "page",
            "in": "query",
            "description": "조회할 페이지 번호 (0부터 시작)",
            "schema": {
              "type": "integer",
              "default": 0
            }
          },
          {
            "name": "size",
            "in": "query",
            "description": "한 페이지에 포함할 항목 수",
            "schema": {
              "type": "integer",
              "default": 20
            }
          },
          {
            "name": "sort",
            "in": "query",
            "description": "정렬 기준 (예: expenseDate,desc)",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpensePageableDtoV2"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v2/expenses/not-submitted": {
      "get": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "미제출 지출내역 목록 조회",
        "description": "아직 용도가 지정되지 않은(미제출, NOT_SUBMITTED) 지출내역 목록을 조회합니다.\n\n- 조회 범위는 **오늘 기준 직전 약 60일(60일 전 ~ 익일)** 입니다. 별도의 기간 파라미터는 받지 않습니다.\n- 용도를 지정해 제출하면 해당 내역은 이 목록에서 빠집니다. (용도 지정 = 제출)",
        "operationId": "getNotSubmittedExpenses",
        "parameters": [
          {
            "name": "page",
            "in": "query",
            "description": "조회할 페이지 번호 (0부터 시작)",
            "schema": {
              "type": "integer",
              "default": 0
            }
          },
          {
            "name": "size",
            "in": "query",
            "description": "한 페이지에 포함할 항목 수",
            "schema": {
              "type": "integer",
              "default": 20
            }
          },
          {
            "name": "sort",
            "in": "query",
            "description": "정렬 기준 (예: expenseDate,desc)",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseSimplePageableDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/expense-statements": {
      "get": {
        "tags": [
          "지출내역 V2"
        ],
        "summary": "지출내역 검색·목록 조회",
        "description": "기간·용도·사용자·메모 등의 조건으로 지출내역을 검색합니다.\n\n- 기간(`startDate`~`endDate`), 승인상태, 카드, 사용자, 메모 등으로 필터링할 수 있습니다.\n- `startDate`/`endDate`는 `yyyyMMdd` 형식입니다.\n- `startDate`·`endDate` 모두 **해당 일자를 포함**합니다(inclusive). 예: `endDate=20260601` 이면 `20260601` 당일 사용분까지 포함됩니다.",
        "operationId": "getExpenseStatements",
        "parameters": [
          {
            "name": "criteria",
            "in": "query",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/ExpenseSearchCriteriaV2"
            }
          },
          {
            "name": "page",
            "in": "query",
            "description": "조회할 페이지 번호 (0부터 시작)",
            "schema": {
              "type": "integer",
              "default": 0
            }
          },
          {
            "name": "size",
            "in": "query",
            "description": "한 페이지에 포함할 항목 수 (최대 100, 초과 시 400 오류)",
            "schema": {
              "maximum": 100,
              "type": "integer",
              "default": 20
            }
          },
          {
            "name": "sort",
            "in": "query",
            "description": "정렬 기준 (예: expenseDate,desc)",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpensePageableDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v2/cards": {
      "get": {
        "tags": [
          "카드 V2"
        ],
        "summary": "카드 목록 조회",
        "description": "법인이 보유한 전체 카드 목록을 조회합니다.\n\n- 각 카드의 한도(`limitAmount` / `usedAmount` / `remainAmount`)와 공동소지자(사용자·부서) 정보를 함께 반환합니다.\n- 결제 이력이 없는 카드(발급만 된 카드 등)도 포함됩니다.",
        "operationId": "getCards",
        "parameters": [
          {
            "name": "page",
            "in": "query",
            "description": "조회할 페이지 번호 (0부터 시작)",
            "schema": {
              "type": "integer",
              "default": 0
            }
          },
          {
            "name": "size",
            "in": "query",
            "description": "한 페이지에 포함할 항목 수 (최대 100, 초과 시 400 오류)",
            "schema": {
              "maximum": 100,
              "type": "integer",
              "default": 20
            }
          },
          {
            "name": "sort",
            "in": "query",
            "description": "정렬 기준 (예: limitAmount,desc)",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseCardPageableDtoV2"
                }
              }
            }
          }
        }
      }
    },
    "/v1/purposes": {
      "get": {
        "tags": [
          "용도"
        ],
        "summary": "사용용도 목록 조회",
        "description": "[Deprecated] V2(`GET /v2/purposes`) 사용을 권장합니다. V2는 세부입력항목(`requirements`) 등 상세 정보를 더 제공합니다.\n\n법인에 설정된 사용용도 목록을 조회합니다. `isActivated`: `true`이면 활성, `false`이면 비활성, 생략 시 전체를 조회합니다.",
        "operationId": "getPurposes_1",
        "parameters": [
          {
            "name": "isActivated",
            "in": "query",
            "required": false,
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseListPurposeDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/purposes/{purposeId}/requirements": {
      "get": {
        "tags": [
          "용도"
        ],
        "summary": "용도 세부입력항목 선택지 목록 조회",
        "description": "[Deprecated] V2(`GET /v2/purposes/{purposeId}/requirements/{requirementId}`) 사용을 권장합니다.\n\n특정 용도의 세부입력항목 선택지 값 목록을 조회합니다. SELECT / SELECT_MULTI(선택형) 타입에서만 값이 반환됩니다. V1은 항목 ID를 지정하지 않아 해당 용도의 첫 번째 선택형 항목을 자동으로 사용합니다.",
        "operationId": "getPurposeRequirements_1",
        "parameters": [
          {
            "name": "purposeId",
            "in": "path",
            "required": true,
            "schema": {
              "type": "integer",
              "format": "int64"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponsePurposeRequirementResDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/members": {
      "get": {
        "tags": [
          "멤버"
        ],
        "summary": "법인 소속 멤버 목록 조회",
        "description": "API Key에 매핑된 법인에 소속된 멤버(사용자) 목록을 조회합니다.\n\n- 삭제(`DELETED`) 상태를 제외한 멤버를 **이름 오름차순**으로 정렬해 반환합니다.\n- 각 멤버의 부서, 직급, 권한(`role`), 상태(`status`) 정보를 포함합니다.",
        "operationId": "getUsersByCorp",
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseListUserSearchResDto"
                }
              }
            }
          }
        }
      }
    },
    "/v1/expenses": {
      "get": {
        "tags": [
          "지출내역"
        ],
        "summary": "지출내역 검색·목록 조회",
        "description": "[Deprecated] V2(`GET /v2/expense-statements`) 사용을 권장합니다.\n\n용도·사용자·메모 키워드로 지출내역을 검색합니다. 키워드가 없으면 해당 기간 내 전체 목록을 반환합니다.\n\n- 종료일(`endDate`)을 입력받지 않으며, 조회 기간은 `startDate` 기준 약 1개월로 **자동 보정**됩니다. 종료일 당일은 **제외**(exclusive)됩니다.",
        "operationId": "searchExpenseHistory_1",
        "parameters": [
          {
            "name": "criteria",
            "in": "query",
            "required": true,
            "schema": {
              "$ref": "#/components/schemas/ExpenseSearchCriteria"
            }
          },
          {
            "name": "page",
            "in": "query",
            "description": "조회할 페이지 번호 (0부터 시작)",
            "schema": {
              "type": "integer",
              "default": 0
            }
          },
          {
            "name": "size",
            "in": "query",
            "description": "한 페이지에 포함할 항목 수",
            "schema": {
              "type": "integer",
              "default": 20
            }
          },
          {
            "name": "sort",
            "in": "query",
            "description": "정렬 기준 (예: expenseDate,desc)",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpensePageableDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    },
    "/v1/expenses/not-submitted": {
      "get": {
        "tags": [
          "지출내역"
        ],
        "summary": "미제출 지출내역 목록 조회",
        "description": "[Deprecated] V2(`GET /v2/expenses/not-submitted`) 사용을 권장합니다.\n\n아직 용도가 지정되지 않은(미제출) 지출내역 목록을 조회합니다. 조회 범위는 오늘 기준 직전 약 60일(60일 전 ~ 익일)입니다.",
        "operationId": "getNotSubmittedExpenses_1",
        "parameters": [
          {
            "name": "page",
            "in": "query",
            "description": "조회할 페이지 번호 (0부터 시작)",
            "schema": {
              "type": "integer",
              "default": 0
            }
          },
          {
            "name": "size",
            "in": "query",
            "description": "한 페이지에 포함할 항목 수",
            "schema": {
              "type": "integer",
              "default": 20
            }
          },
          {
            "name": "sort",
            "in": "query",
            "description": "정렬 기준 (예: expenseDate,desc)",
            "schema": {
              "type": "string"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "content": {
              "application/json": {
                "schema": {
                  "$ref": "#/components/schemas/GowidResponseExpenseSimplePageableDto"
                }
              }
            }
          }
        },
        "deprecated": true
      }
    }
  },
  "components": {
    "schemas": {
      "CardDtoV2": {
        "type": "object",
        "properties": {
          "cardId": {
            "type": "integer",
            "description": "카드 id",
            "format": "int64"
          },
          "companyCode": {
            "type": "string",
            "description": "카드사 코드"
          },
          "cardStatus": {
            "type": "string",
            "description": "카드 상태"
          },
          "cardAlias": {
            "type": "string",
            "description": "카드 별칭"
          },
          "managementNumber": {
            "type": "string",
            "description": "관리번호"
          },
          "maskedCardNumber": {
            "type": "string",
            "description": "마스킹된 카드번호. 앞 6자리와 뒤 4자리만 노출됩니다.",
            "example": "552576******1234"
          },
          "encryptedCardNumber": {
            "type": "string",
            "description": "암호화된 카드번호. 암호화 관련 내용은 openapi@gowid.com 으로 문의 부탁드립니다."
          },
          "limitAmount": {
            "type": "integer",
            "description": "카드 한도",
            "format": "int64"
          },
          "usedAmount": {
            "type": "integer",
            "description": "사용 금액",
            "format": "int64"
          },
          "remainAmount": {
            "type": "integer",
            "description": "잔여 한도",
            "format": "int64"
          },
          "cardUser": {
            "$ref": "#/components/schemas/CardUserDtoV2"
          },
          "userHolders": {
            "type": "array",
            "description": "공동 소지자 (사용자)",
            "items": {
              "$ref": "#/components/schemas/CardHolderUserDtoV2"
            }
          },
          "departmentHolders": {
            "type": "array",
            "description": "공동 소지자 (부서)",
            "items": {
              "$ref": "#/components/schemas/CardHolderDepartmentDtoV2"
            }
          }
        }
      },
      "CardHolderDepartmentDtoV2": {
        "type": "object",
        "properties": {
          "departmentId": {
            "type": "integer",
            "description": "부서 id",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "부서명"
          }
        }
      },
      "CardHolderUserDtoV2": {
        "type": "object",
        "properties": {
          "userId": {
            "type": "integer",
            "description": "사용자 id",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "사용자 이름"
          },
          "departmentName": {
            "type": "string",
            "description": "부서명"
          }
        }
      },
      "CardPageableDtoV2": {
        "type": "object",
        "properties": {
          "totalPages": {
            "type": "integer",
            "description": "전체 페이지 수",
            "format": "int32"
          },
          "totalElements": {
            "type": "integer",
            "description": "전체 요소 수",
            "format": "int32"
          },
          "last": {
            "type": "boolean",
            "description": "마지막 페이지 여부"
          },
          "content": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/CardDtoV2"
            }
          }
        }
      },
      "CardUserDtoV2": {
        "type": "object",
        "properties": {
          "userId": {
            "type": "integer",
            "description": "사용자 id",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "사용자 이름"
          },
          "departmentName": {
            "type": "string",
            "description": "부서명"
          }
        },
        "description": "대표 소지자"
      },
      "CardVo": {
        "type": "object",
        "properties": {
          "cardNumber": {
            "type": "string",
            "description": "카드번호"
          },
          "encryptedCardNumber": {
            "type": "string",
            "description": "암호화된 카드번호. 암호화 관련 내용은 openapi@gowid.com 으로 문의 부탁드립니다."
          },
          "cardUser": {
            "$ref": "#/components/schemas/UserVo"
          },
          "alias": {
            "type": "string",
            "description": "카드 별칭"
          },
          "limitAmount": {
            "type": "integer",
            "description": "카드 한도",
            "format": "int64"
          },
          "usedAmount": {
            "type": "integer",
            "description": "사용 한도",
            "format": "int64"
          },
          "remainAmount": {
            "type": "integer",
            "description": "잔여 한도",
            "format": "int64"
          },
          "companyCode": {
            "type": "string",
            "description": "카드사 코드. 0305=비씨, 0306=신한, 0311=롯데"
          },
          "namedYn": {
            "type": "string",
            "description": "기명카드 여부 (\"Y\" 또는 \"N\")."
          },
          "cardName": {
            "type": "string",
            "description": "카드명"
          },
          "cardType": {
            "type": "string",
            "description": "카드 구분 (카드사 코드값)."
          },
          "userNm": {
            "type": "string",
            "description": "카드 소지자 성명 (카드사 원문 데이터 기준)."
          },
          "duplicationStatus": {
            "type": "string",
            "description": "중복카드 보정 상태",
            "enum": [
              "NORMAL",
              "NOT_PROC",
              "PROCESSING",
              "COMPLETED",
              "DUP_FAILED",
              "WRONG_FAILED",
              "SYS_FAILED"
            ]
          },
          "invalid": {
            "type": "boolean"
          }
        },
        "description": "카드 정보"
      },
      "CategoryVo": {
        "type": "object",
        "properties": {
          "categoryId": {
            "type": "integer",
            "description": "카테고리 id",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "카테고리명"
          }
        },
        "description": "카테고리 정보"
      },
      "CommentReqDto": {
        "required": [
          "comment"
        ],
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "댓글을 작성할 지출 ID",
            "format": "int64"
          },
          "comment": {
            "type": "string",
            "description": "댓글 내용",
            "example": "영수증 확인 완료"
          }
        }
      },
      "CommentResDto": {
        "type": "object",
        "properties": {
          "author": {
            "type": "string",
            "description": "댓글 작성자 이름"
          },
          "department": {
            "type": "string",
            "description": "작성자 소속 부서명"
          },
          "content": {
            "type": "string",
            "description": "댓글 내용"
          },
          "createdAt": {
            "type": "string",
            "description": "댓글 작성 일시 (ISO-8601)",
            "format": "date-time"
          }
        }
      },
      "CommentVo": {
        "type": "object",
        "properties": {
          "author": {
            "$ref": "#/components/schemas/UserVo"
          },
          "content": {
            "type": "string",
            "description": "댓글 내용"
          },
          "createdAt": {
            "type": "string",
            "description": "작성시간",
            "format": "date-time"
          }
        }
      },
      "DepartmentResDto": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer",
            "description": "부서 ID",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "부서명",
            "example": "재무팀"
          }
        },
        "description": "소속 부서 정보"
      },
      "ExpenseApproveReqDto": {
        "required": [
          "approvalStatus"
        ],
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 id",
            "format": "int64"
          },
          "approvalStatus": {
            "type": "string",
            "description": "변경할 승인 상태. APPROVED(승인), REJECTED(반려). 부분승인은 직접 지정하지 않고 APPROVED + approvedAmount 조합으로 자동 처리됨.",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "approvedAmount": {
            "type": "integer",
            "description": "관리자 승인금액 (원화). 전체 금액보다 작게 보내면 부분승인 처리됨.",
            "format": "int64",
            "example": 50000
          },
          "approvedAt": {
            "type": "string",
            "description": "관리자 승인일시 (yyyyMMddHHmmss)",
            "example": "20260108150121"
          }
        }
      },
      "ExpenseApproveReqDtoV2": {
        "required": [
          "approvalStatus"
        ],
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 id",
            "format": "int64"
          },
          "approvalStatus": {
            "type": "string",
            "description": "관리자 승인상태 (미제출, 제출, 승인, 부분승인, 반려)",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "approvedAmount": {
            "type": "integer",
            "description": "관리자 승인금액",
            "format": "int64",
            "example": 50000
          },
          "approvedAt": {
            "type": "string",
            "description": "관리자 승인일시 (yyyyMMddHHmmss)",
            "example": "20260108150121"
          }
        }
      },
      "ExpenseDeductionResDto": {
        "type": "object",
        "properties": {
          "isExpenseDeductible": {
            "type": "boolean",
            "description": "법인 설정 - 이용내역별 공제 가능 설정 여부"
          },
          "isDeducted": {
            "type": "boolean",
            "description": "이용내역 공제 여부"
          }
        },
        "description": "지출 공제 정보"
      },
      "ExpenseDetailResDto": {
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 id",
            "format": "int64"
          },
          "cardApprovalNumber": {
            "type": "string",
            "description": "카드 승인 번호"
          },
          "expenseDate": {
            "type": "string",
            "description": "사용연월일 (yyyyMMdd)",
            "example": "20260108"
          },
          "expenseTime": {
            "type": "string",
            "description": "사용일시 (HHmmss)",
            "example": "103024"
          },
          "card": {
            "$ref": "#/components/schemas/CardVo"
          },
          "user": {
            "$ref": "#/components/schemas/UserVo"
          },
          "useAmount": {
            "type": "number",
            "description": "현지 사용금액",
            "example": 35.0
          },
          "currency": {
            "type": "string",
            "description": "사용 화폐",
            "example": "USD"
          },
          "krwAmount": {
            "type": "integer",
            "description": "원화 환산금액",
            "format": "int64",
            "example": 50732
          },
          "approvalStatus": {
            "type": "string",
            "description": "관리자 승인상태 (미제출, 제출, 승인, 부분승인, 반려)",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "approvedAmount": {
            "type": "integer",
            "description": "관리자 승인금액",
            "format": "int64",
            "example": 50000
          },
          "approvedAt": {
            "type": "string",
            "description": "관리자 승인일시 (yyyyMMddHHmmss)",
            "example": "20260108150121"
          },
          "approvedBy": {
            "type": "string",
            "description": "승인자"
          },
          "comments": {
            "type": "array",
            "description": "댓글 목록",
            "items": {
              "$ref": "#/components/schemas/CommentVo"
            }
          },
          "purpose": {
            "$ref": "#/components/schemas/PurposeVo"
          },
          "participants": {
            "type": "array",
            "description": "참석자 목록",
            "items": {
              "$ref": "#/components/schemas/UserVo"
            }
          },
          "expenseExternalUsers": {
            "type": "array",
            "description": "외부 참석자 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseExternalUserVo"
            }
          },
          "storeName": {
            "type": "string",
            "description": "사용처",
            "example": "UNITED"
          },
          "storeAddress": {
            "type": "string",
            "description": "사용처 주소"
          },
          "storeRegistrationNumber": {
            "type": "string",
            "description": "사용처 사업자등록번호. 해외결제건이거나 조회 불가능한 경우 null이 반환됩니다."
          },
          "memo": {
            "type": "string",
            "description": "메모"
          },
          "evidenceList": {
            "type": "array",
            "description": "영수증 첨부파일 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseEvidenceVo"
            }
          },
          "companyCode": {
            "type": "string",
            "description": "카드사 코드. 0305=비씨, 0306=신한, 0311=롯데"
          },
          "commentCount": {
            "type": "integer",
            "description": "댓글 수",
            "format": "int32"
          },
          "purposeRequirementItem": {
            "type": "string",
            "description": "용도 필수입력항목 제목"
          },
          "purposeRequirementItemType": {
            "type": "string",
            "description": "용도 필수입력항목 타입",
            "enum": [
              "TEXT",
              "SELECT",
              "SELECT_MULTI"
            ]
          },
          "purposeRequirementValue": {
            "type": "string",
            "description": "용도 필수입력항목 응답"
          },
          "expenseDeductionResDto": {
            "$ref": "#/components/schemas/ExpenseDeductionResDto"
          },
          "isDomestic": {
            "type": "boolean",
            "description": "국내결제 여부"
          },
          "isPurposeRequirementInputValue": {
            "type": "boolean",
            "description": "용도 필수입력항목 직접입력 여부"
          }
        }
      },
      "ExpenseDetailResDtoV2": {
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 id",
            "format": "int64"
          },
          "cardApprovalNumber": {
            "type": "string",
            "description": "카드 승인 번호"
          },
          "expenseType": {
            "type": "string",
            "description": "카드사 거래상태",
            "enum": [
              "NOT_DEFINE",
              "APPROVAL",
              "PURCHASE",
              "BILLING",
              "CANCEL_APPROVAL",
              "PARTIAL_CANCEL_APPROVAL",
              "CANCEL_PURCHASE",
              "PARTIAL_CANCEL_PURCHASE"
            ]
          },
          "expenseDate": {
            "type": "string",
            "description": "사용연월일 (yyyyMMdd)",
            "example": "20260108"
          },
          "expenseTime": {
            "type": "string",
            "description": "사용일시 (HHmmss)",
            "example": "103024"
          },
          "card": {
            "$ref": "#/components/schemas/CardVo"
          },
          "user": {
            "$ref": "#/components/schemas/UserVo"
          },
          "useAmount": {
            "type": "number",
            "description": "현지 사용금액",
            "example": 35.0
          },
          "currency": {
            "type": "string",
            "description": "사용 화폐",
            "example": "USD"
          },
          "krwAmount": {
            "type": "integer",
            "description": "원화 환산금액",
            "format": "int64",
            "example": 50732
          },
          "approvalStatus": {
            "type": "string",
            "description": "관리자 승인상태 (미제출, 제출, 승인, 부분승인, 반려)",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "approvedAmount": {
            "type": "integer",
            "description": "관리자 승인금액",
            "format": "int64",
            "example": 50000
          },
          "approvedAt": {
            "type": "string",
            "description": "관리자 승인일시 (yyyyMMddHHmmss)",
            "example": "20260108150121"
          },
          "approvedBy": {
            "type": "string",
            "description": "승인자"
          },
          "comments": {
            "type": "array",
            "description": "댓글 목록",
            "items": {
              "$ref": "#/components/schemas/CommentVo"
            }
          },
          "purpose": {
            "$ref": "#/components/schemas/PurposeVo"
          },
          "participants": {
            "type": "array",
            "description": "참석자 목록",
            "items": {
              "$ref": "#/components/schemas/UserVo"
            }
          },
          "expenseExternalUsers": {
            "type": "array",
            "description": "외부 참석자 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseExternalUserVo"
            }
          },
          "storeName": {
            "type": "string",
            "description": "사용처",
            "example": "UNITED"
          },
          "storeAddress": {
            "type": "string",
            "description": "사용처 주소"
          },
          "storeRegistrationNumber": {
            "type": "string",
            "description": "사용처 사업자등록번호. 해외결제건이거나 조회 불가능한 경우 null이 반환됩니다."
          },
          "memo": {
            "type": "string",
            "description": "메모"
          },
          "evidenceList": {
            "type": "array",
            "description": "영수증 첨부파일 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseEvidenceVo"
            }
          },
          "companyCode": {
            "type": "string",
            "description": "카드사 코드. 0305=비씨, 0306=신한, 0311=롯데"
          },
          "commentCount": {
            "type": "integer",
            "description": "댓글 수",
            "format": "int32"
          },
          "expenseDeductionResDto": {
            "$ref": "#/components/schemas/ExpenseDeductionResDto"
          },
          "purposeRequirementAnswers": {
            "type": "array",
            "description": "세부입력항목 응답 목록",
            "items": {
              "$ref": "#/components/schemas/PurposeRequirementAnswerDtoV2"
            }
          },
          "isDomestic": {
            "type": "boolean",
            "description": "국내결제 여부"
          }
        }
      },
      "ExpenseDto": {
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 ID",
            "format": "int64"
          },
          "expenseDate": {
            "type": "string",
            "description": "결제 연월일 (yyyyMMdd)",
            "example": "20260108"
          },
          "expenseTime": {
            "type": "string",
            "description": "결제 시각 (HHmmss)",
            "example": "103024"
          },
          "useAmount": {
            "type": "number",
            "description": "현지 결제 금액",
            "example": 35.0
          },
          "currency": {
            "type": "string",
            "description": "결제 통화 코드 (ISO 4217)",
            "example": "KRW"
          },
          "krwAmount": {
            "type": "integer",
            "description": "원화 환산 금액",
            "format": "int64",
            "example": 50000
          },
          "approvedAmount": {
            "type": "integer",
            "description": "관리자 승인 금액 (원화)",
            "format": "int64",
            "example": 50000
          },
          "approvedAt": {
            "type": "string",
            "description": "관리자 승인 일시 (yyyyMMddHHmmss)",
            "example": "20260108150121"
          },
          "approvalStatus": {
            "type": "string",
            "description": "관리자 승인 상태 (NOT_SUBMITTED=미제출, SUBMITTED=승인대기, APPROVED=승인, PARTIAL_APPROVED=부분승인, REJECTED=반려)",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "purpose": {
            "$ref": "#/components/schemas/PurposeSimpleDto"
          },
          "cardAlias": {
            "type": "string",
            "description": "카드 별칭"
          },
          "cardUserName": {
            "type": "string",
            "description": "카드 소지자 이름"
          },
          "shortCardNumber": {
            "type": "string",
            "description": "카드번호 끝 4자리",
            "example": "1234"
          },
          "storeName": {
            "type": "string",
            "description": "가맹점(사용처) 이름"
          },
          "storeAddress": {
            "type": "string",
            "description": "가맹점 주소"
          },
          "memo": {
            "type": "string",
            "description": "결제자 입력 메모"
          },
          "commentCount": {
            "type": "integer",
            "description": "댓글 수",
            "format": "int32"
          },
          "evidenceCount": {
            "type": "integer",
            "description": "첨부 영수증 파일 수",
            "format": "int32"
          },
          "participantCount": {
            "type": "integer",
            "description": "전체 참석자 수 (내부 + 외부 참석자)",
            "format": "int32"
          },
          "representativeParticipant": {
            "type": "string",
            "description": "대표 참석자 표시 문자열. 카드 소지자가 참석자에 포함되면 카드 소지자 이름, 아니면 첫 번째 참석자 이름 기준 (예: '홍길동 외 2명')."
          },
          "participants": {
            "type": "array",
            "description": "내부 참석자 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseParticipantDto"
            }
          },
          "expenseExternalUsers": {
            "type": "array",
            "description": "외부 참석자 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseExternalUserVo"
            }
          },
          "purposeRequirementItem": {
            "type": "string",
            "description": "세부입력항목 이름 (V1 레거시, 단일 항목)"
          },
          "purposeRequirementItemType": {
            "type": "string",
            "description": "세부입력항목 응답 유형 (TEXT=텍스트, SELECT=단일선택, SELECT_MULTI=복수선택)",
            "enum": [
              "TEXT",
              "SELECT",
              "SELECT_MULTI"
            ]
          },
          "purposeRequirementValue": {
            "type": "string",
            "description": "세부입력항목 입력값"
          },
          "isPurposeRequirementInputValue": {
            "type": "boolean",
            "description": "세부입력항목 직접 입력 허용 여부"
          }
        }
      },
      "ExpenseDtoV2": {
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 id",
            "format": "int64"
          },
          "cardApprovalNumber": {
            "type": "string",
            "description": "카드 승인 번호"
          },
          "expenseDate": {
            "type": "string",
            "description": "사용연월일 (yyyyMMdd)",
            "example": "20260108"
          },
          "expenseTime": {
            "type": "string",
            "description": "사용일시 (HHmmss)",
            "example": "103024"
          },
          "expenseType": {
            "type": "string",
            "description": "카드사 거래상태",
            "enum": [
              "NOT_DEFINE",
              "APPROVAL",
              "PURCHASE",
              "BILLING",
              "CANCEL_APPROVAL",
              "PARTIAL_CANCEL_APPROVAL",
              "CANCEL_PURCHASE",
              "PARTIAL_CANCEL_PURCHASE"
            ]
          },
          "useAmount": {
            "type": "number",
            "description": "현지 사용금액",
            "example": 35.0
          },
          "currency": {
            "type": "string",
            "description": "사용 화폐",
            "example": "USD"
          },
          "krwAmount": {
            "type": "integer",
            "description": "원화 환산금액",
            "format": "int64",
            "example": 50732
          },
          "approvedAmount": {
            "type": "integer",
            "description": "관리자 승인금액",
            "format": "int64",
            "example": 50000
          },
          "approvedAt": {
            "type": "string",
            "description": "관리자 승인일시 (yyyyMMddHHmmss)",
            "example": "20260108150121"
          },
          "approvalStatus": {
            "type": "string",
            "description": "관리자 승인상태 (미제출, 제출, 승인, 부분승인, 반려)",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "purpose": {
            "$ref": "#/components/schemas/PurposeSimpleDtoV2"
          },
          "cardAlias": {
            "type": "string",
            "description": "카드 별칭"
          },
          "cardUserName": {
            "type": "string",
            "description": "카드 소지자 이름"
          },
          "shortCardNumber": {
            "type": "string",
            "description": "카드번호 끝 4자리"
          },
          "encryptedCardNumber": {
            "type": "string",
            "description": "암호화된 카드번호. 암호화 관련 내용은 openapi@gowid.com 으로 문의 부탁드립니다."
          },
          "storeName": {
            "type": "string",
            "description": "사용처"
          },
          "storeAddress": {
            "type": "string",
            "description": "사용처 주소"
          },
          "memo": {
            "type": "string",
            "description": "메모"
          },
          "commentCount": {
            "type": "integer",
            "description": "댓글 수",
            "format": "int32"
          },
          "evidenceCount": {
            "type": "integer",
            "description": "영수증 첨부파일 수",
            "format": "int32"
          },
          "participantCount": {
            "type": "integer",
            "description": "참석자 수 (외부인원 포함)",
            "format": "int32"
          },
          "representativeParticipant": {
            "type": "string",
            "description": "대표 참석자"
          },
          "participants": {
            "type": "array",
            "description": "참석자",
            "items": {
              "$ref": "#/components/schemas/ExpenseParticipantDtoV2"
            }
          },
          "expenseExternalUsers": {
            "type": "array",
            "description": "외부 참석자",
            "items": {
              "$ref": "#/components/schemas/ExpenseExternalUserVo"
            }
          },
          "purposeRequirementAnswers": {
            "type": "array",
            "description": "세부입력항목 응답 목록",
            "items": {
              "$ref": "#/components/schemas/PurposeRequirementAnswerDtoV2"
            }
          }
        }
      },
      "ExpenseEvidenceVo": {
        "type": "object",
        "properties": {
          "evidenceId": {
            "type": "integer",
            "description": "첨부파일 ID",
            "format": "int64",
            "example": 1001
          },
          "fileName": {
            "type": "string",
            "description": "저장된 파일명",
            "example": "receipt_20260108.jpg"
          },
          "mimeType": {
            "type": "string",
            "description": "파일 MIME 타입",
            "example": "image/jpeg"
          },
          "signedUrl": {
            "type": "string",
            "description": "첨부 영수증 다운로드용 Signed URL. 일정 시간 후 만료되므로 즉시 사용 권장.",
            "example": "https://storage.gowid.com/..."
          }
        }
      },
      "ExpenseExternalUserDto": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string",
            "description": "외부 참석자 이름",
            "example": "홍길동"
          },
          "company": {
            "type": "string",
            "description": "외부 참석자 소속 회사",
            "example": "삼성전자"
          }
        }
      },
      "ExpenseExternalUserDtoV2": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string",
            "description": "외부 참석자 이름"
          },
          "company": {
            "type": "string",
            "description": "외부 참석자 소속 회사"
          }
        }
      },
      "ExpenseExternalUserVo": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string",
            "description": "이름"
          },
          "company": {
            "type": "string",
            "description": "회사명"
          }
        }
      },
      "ExpenseMemoReqDto": {
        "required": [
          "memo"
        ],
        "type": "object",
        "properties": {
          "memo": {
            "type": "string",
            "description": "메모",
            "example": "팀 회식"
          }
        }
      },
      "ExpenseMemoReqDtoV2": {
        "required": [
          "memo"
        ],
        "type": "object",
        "properties": {
          "memo": {
            "type": "string",
            "description": "메모"
          }
        }
      },
      "ExpensePageableDto": {
        "type": "object",
        "properties": {
          "totalPages": {
            "type": "integer",
            "description": "전체 페이지 수",
            "format": "int32"
          },
          "totalElements": {
            "type": "integer",
            "description": "전체 요소 수",
            "format": "int32"
          },
          "last": {
            "type": "boolean",
            "description": "마지막 페이지 여부"
          },
          "content": {
            "type": "array",
            "description": "지출 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseDto"
            }
          }
        }
      },
      "ExpensePageableDtoV2": {
        "type": "object",
        "properties": {
          "totalPages": {
            "type": "integer",
            "description": "전체 페이지 수",
            "format": "int32"
          },
          "totalElements": {
            "type": "integer",
            "description": "전체 요소 수",
            "format": "int32"
          },
          "last": {
            "type": "boolean",
            "description": "마지막 페이지 여부"
          },
          "content": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/ExpenseDtoV2"
            }
          }
        }
      },
      "ExpenseParticipantDto": {
        "type": "object",
        "properties": {
          "userId": {
            "type": "integer",
            "description": "참석자 ID",
            "format": "int64"
          },
          "userName": {
            "type": "string",
            "description": "참석자명"
          }
        }
      },
      "ExpenseParticipantDtoV2": {
        "type": "object",
        "properties": {
          "userId": {
            "type": "integer",
            "description": "참석자 ID",
            "format": "int64"
          },
          "userName": {
            "type": "string",
            "description": "참석자명"
          }
        }
      },
      "ExpenseParticipantsUpdateRequestDto": {
        "required": [
          "externalUsers",
          "participantIds"
        ],
        "type": "object",
        "properties": {
          "participantIds": {
            "type": "array",
            "description": "내부 참석자 ID 목록",
            "items": {
              "type": "integer",
              "description": "내부 참석자 ID 목록",
              "format": "int64"
            }
          },
          "externalUsers": {
            "type": "array",
            "description": "외부 참석자 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseExternalUserDto"
            }
          }
        }
      },
      "ExpenseParticipantsUpdateRequestDtoV2": {
        "required": [
          "externalUsers",
          "participantIds"
        ],
        "type": "object",
        "properties": {
          "participantIds": {
            "type": "array",
            "description": "내부 참석자 ID 목록",
            "items": {
              "type": "integer",
              "description": "내부 참석자 ID 목록",
              "format": "int64"
            }
          },
          "externalUsers": {
            "type": "array",
            "description": "외부 참석자 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseExternalUserDtoV2"
            }
          }
        }
      },
      "ExpensePurposeBulkUpdateRequestDtoV2": {
        "required": [
          "expenseIds",
          "purposeId"
        ],
        "type": "object",
        "properties": {
          "expenseIds": {
            "maxItems": 2147483647,
            "minItems": 1,
            "type": "array",
            "description": "용도를 변경할 지출들의 id",
            "items": {
              "type": "integer",
              "description": "용도를 변경할 지출들의 id",
              "format": "int64"
            }
          },
          "purposeId": {
            "type": "integer",
            "description": "설정할 용도 id",
            "format": "int64"
          },
          "purposeRequirementAnswerMap": {
            "type": "object",
            "additionalProperties": {
              "type": "array",
              "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요.",
              "items": {
                "type": "string",
                "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요."
              }
            },
            "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요."
          }
        }
      },
      "ExpensePurposeUpdateRequestDto": {
        "required": [
          "purposeId"
        ],
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "용도를 변경할 지출 ID",
            "format": "int64"
          },
          "purposeId": {
            "type": "integer",
            "description": "설정할 용도 id",
            "format": "int64"
          },
          "purposeRequirementItem": {
            "type": "string",
            "description": "세부입력항목 이름 (V1 단일 항목 방식. V2는 purposeRequirementAnswerMap 사용)"
          },
          "purposeRequirementItemType": {
            "type": "string",
            "description": "세부입력항목 응답 유형 (TEXT=텍스트, SELECT=단일선택)",
            "enum": [
              "TEXT",
              "SELECT",
              "SELECT_MULTI"
            ]
          },
          "purposeRequirementValue": {
            "maxLength": 40,
            "minLength": 0,
            "type": "string",
            "description": "세부입력항목 응답 값 (최대 40자)"
          },
          "purposeRequirementInputValue": {
            "type": "boolean"
          },
          "isPurposeRequirementInputValue": {
            "type": "boolean",
            "description": "세부입력항목 직접입력 허용 여부. type=SELECT일 때 선택지 외 직접 입력이 허용된 경우 true."
          }
        }
      },
      "ExpensePurposeUpdateRequestDtoV2": {
        "required": [
          "purposeId"
        ],
        "type": "object",
        "properties": {
          "purposeId": {
            "type": "integer",
            "description": "설정할 용도 id",
            "format": "int64"
          },
          "purposeRequirementAnswerMap": {
            "type": "object",
            "additionalProperties": {
              "type": "array",
              "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요.",
              "items": {
                "type": "string",
                "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요."
              }
            },
            "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요."
          }
        }
      },
      "ExpenseSearchCriteria": {
        "type": "object",
        "properties": {
          "memo": {
            "type": "string",
            "description": "메모 키워드 검색 (부분 일치)"
          },
          "purposeName": {
            "type": "string",
            "description": "용도 이름으로 검색 (부분 일치)"
          },
          "userName": {
            "type": "string",
            "description": "카드 소지자 이름으로 검색 (부분 일치)"
          },
          "startDate": {
            "type": "string",
            "description": "검색 시작일 (yyyyMMdd)",
            "example": "20260101"
          },
          "approvalState": {
            "type": "string",
            "description": "관리자 승인 상태 필터 (NOT_SUBMITTED, SUBMITTED, APPROVED, PARTIAL_APPROVED, REJECTED)",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          }
        }
      },
      "ExpenseSearchCriteriaV2": {
        "type": "object",
        "properties": {
          "memo": {
            "type": "string",
            "description": "메모"
          },
          "purposeName": {
            "type": "string",
            "description": "용도 이름"
          },
          "userName": {
            "type": "string",
            "description": "사용자 이름"
          },
          "startDate": {
            "type": "string",
            "description": "시작일"
          },
          "endDate": {
            "type": "string",
            "description": "종료일"
          },
          "approvalState": {
            "type": "string",
            "description": "관리자 승인 상태",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "encryptedCardNumber": {
            "type": "string",
            "description": "암호화된 카드번호. 암호화 관련 내용은 openapi@gowid.com 으로 문의 부탁드립니다."
          }
        }
      },
      "ExpenseSimpleDto": {
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 ID",
            "format": "int64"
          },
          "expenseDate": {
            "type": "string",
            "description": "결제 연월일 (yyyyMMdd)",
            "example": "20260108"
          },
          "expenseTime": {
            "type": "string",
            "description": "결제 시각 (HHmmss)",
            "example": "103024"
          },
          "useAmount": {
            "type": "number",
            "description": "현지 결제 금액",
            "example": 35.0
          },
          "currency": {
            "type": "string",
            "description": "결제 통화 코드 (ISO 4217)",
            "example": "KRW"
          },
          "krwAmount": {
            "type": "integer",
            "description": "원화 환산 금액",
            "format": "int64",
            "example": 50000
          },
          "storeName": {
            "type": "string",
            "description": "가맹점(사용처) 이름"
          },
          "approvalStatus": {
            "type": "string",
            "description": "관리자 승인 상태 (NOT_SUBMITTED=미제출, SUBMITTED=승인대기, APPROVED=승인, PARTIAL_APPROVED=부분승인, REJECTED=반려)",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "cardAlias": {
            "type": "string",
            "description": "카드 별칭"
          },
          "shortCardNumber": {
            "type": "string",
            "description": "카드번호 끝 4자리",
            "example": "1234"
          },
          "recommendedPurposeList": {
            "type": "array",
            "description": "시스템 추천 용도 목록. 용도가 아직 지정되지 않은(미제출) 내역에 대해, 최근 60일 내 동일 가맹점에서 사용된 용도 또는 활성 용도 상위 항목을 추천.",
            "items": {
              "$ref": "#/components/schemas/PurposeDto"
            }
          }
        }
      },
      "ExpenseSimpleDtoV2": {
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 id",
            "format": "int64"
          },
          "expenseDate": {
            "type": "string",
            "description": "사용연월일 (yyyyMMdd)",
            "example": "20260108"
          },
          "expenseTime": {
            "type": "string",
            "description": "사용일시 (HHmmss)",
            "example": "103024"
          },
          "useAmount": {
            "type": "number",
            "description": "현지 사용금액",
            "example": 35.0
          },
          "currency": {
            "type": "string",
            "description": "사용 화폐",
            "example": "USD"
          },
          "krwAmount": {
            "type": "integer",
            "description": "원화 환산금액",
            "format": "int64",
            "example": 50732
          },
          "storeName": {
            "type": "string",
            "description": "사용처",
            "example": "UNITED"
          },
          "approvalStatus": {
            "type": "string",
            "description": "관리자 승인상태 (미제출, 제출, 승인, 부분승인, 반려)",
            "enum": [
              "NOT_SUBMITTED",
              "SUBMITTED",
              "APPROVED",
              "PARTIAL_APPROVED",
              "REJECTED"
            ]
          },
          "cardAlias": {
            "type": "string",
            "description": "카드 별칭"
          },
          "shortCardNumber": {
            "type": "string",
            "description": "카드번호 끝 4자리"
          }
        }
      },
      "ExpenseSimplePageableDto": {
        "type": "object",
        "properties": {
          "totalPages": {
            "type": "integer",
            "description": "전체 페이지 수",
            "format": "int32"
          },
          "totalElements": {
            "type": "integer",
            "description": "전체 요소 수",
            "format": "int32"
          },
          "last": {
            "type": "boolean",
            "description": "마지막 페이지 여부"
          },
          "content": {
            "type": "array",
            "description": "지출내역 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseSimpleDto"
            }
          }
        }
      },
      "ExpenseSimplePageableDtoV2": {
        "type": "object",
        "properties": {
          "totalPages": {
            "type": "integer",
            "description": "전체 페이지 수",
            "format": "int32"
          },
          "totalElements": {
            "type": "integer",
            "description": "전체 요소 수",
            "format": "int32"
          },
          "last": {
            "type": "boolean",
            "description": "마지막 페이지 여부"
          },
          "content": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/ExpenseSimpleDtoV2"
            }
          }
        }
      },
      "ExpenseUpdateReqDto": {
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 id",
            "format": "int64"
          },
          "purposeId": {
            "type": "integer",
            "description": "용도 id",
            "format": "int64"
          },
          "purposeRequirementItem": {
            "type": "string",
            "description": "세부입력항목 이름 (V1 단일 항목 방식. V2는 purposeRequirementAnswerMap 사용)"
          },
          "purposeRequirementItemType": {
            "type": "string",
            "description": "세부입력항목 응답 유형 (TEXT=텍스트, SELECT=단일선택)",
            "enum": [
              "TEXT",
              "SELECT",
              "SELECT_MULTI"
            ]
          },
          "purposeRequirementValue": {
            "maxLength": 40,
            "minLength": 0,
            "type": "string",
            "description": "세부입력항목 응답 값 (최대 40자)"
          },
          "participantIdList": {
            "type": "array",
            "description": "내부 참석자 ID 목록",
            "items": {
              "type": "integer",
              "description": "내부 참석자 ID 목록",
              "format": "int64"
            }
          },
          "externalUserList": {
            "type": "array",
            "description": "외부 참석자 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseExternalUserDto"
            }
          },
          "memo": {
            "type": "string",
            "description": "메모"
          },
          "fileIdList": {
            "type": "array",
            "description": "첨부이미지 ID 목록 (최대 4개). Open API에서는 사용되지 않음.",
            "items": {
              "type": "integer",
              "description": "첨부이미지 ID 목록 (최대 4개). Open API에서는 사용되지 않음.",
              "format": "int64"
            }
          },
          "isIgnoredFiles": {
            "type": "boolean",
            "description": "파일 업데이트 무시 여부. Open API에서는 파일 수정을 지원하지 않으므로 항상 true."
          },
          "isPurposeRequirementInputValue": {
            "type": "boolean",
            "description": "세부입력항목 직접입력 허용 여부. type=SELECT일 때 선택지 외 직접 입력이 허용된 경우 true."
          }
        }
      },
      "ExpenseUpdateReqDtoV2": {
        "type": "object",
        "properties": {
          "expenseId": {
            "type": "integer",
            "description": "지출 id",
            "format": "int64"
          },
          "purposeId": {
            "type": "integer",
            "description": "용도 id",
            "format": "int64"
          },
          "participantIdList": {
            "type": "array",
            "description": "내부 참석자 ID 목록",
            "items": {
              "type": "integer",
              "description": "내부 참석자 ID 목록",
              "format": "int64"
            }
          },
          "externalUserList": {
            "type": "array",
            "description": "외부 참석자 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseExternalUserDtoV2"
            }
          },
          "purposeRequirementAnswerMap": {
            "type": "object",
            "additionalProperties": {
              "type": "array",
              "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요.",
              "items": {
                "type": "string",
                "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요."
              }
            },
            "description": "세부입력항목 답변. key에는 세부입력항목의 id, value에는 답변 목록을 담아주세요."
          },
          "memo": {
            "type": "string",
            "description": "메모"
          },
          "fileIdList": {
            "type": "array",
            "description": "첨부이미지 ID 목록 (최대 4개)",
            "items": {
              "type": "integer",
              "description": "첨부이미지 ID 목록 (최대 4개)",
              "format": "int64"
            }
          },
          "isIgnoredFiles": {
            "type": "boolean",
            "description": "파일 업데이트 무시 여부. 현재 Open API에서는 파일 수정을 제공하지 않습니다."
          }
        }
      },
      "GowidResponseBoolean": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "type": "boolean"
          }
        }
      },
      "GowidResponseCardPageableDtoV2": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/CardPageableDtoV2"
          }
        }
      },
      "GowidResponseCommentResDto": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/CommentResDto"
          }
        }
      },
      "GowidResponseExpenseDetailResDto": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/ExpenseDetailResDto"
          }
        }
      },
      "GowidResponseExpenseDetailResDtoV2": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/ExpenseDetailResDtoV2"
          }
        }
      },
      "GowidResponseExpensePageableDto": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/ExpensePageableDto"
          }
        }
      },
      "GowidResponseExpensePageableDtoV2": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/ExpensePageableDtoV2"
          }
        }
      },
      "GowidResponseExpenseSimplePageableDto": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/ExpenseSimplePageableDto"
          }
        }
      },
      "GowidResponseExpenseSimplePageableDtoV2": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/ExpenseSimplePageableDtoV2"
          }
        }
      },
      "GowidResponseListPurposeDto": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/PurposeDto"
            }
          }
        }
      },
      "GowidResponseListPurposeDtoV2": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/PurposeDtoV2"
            }
          }
        }
      },
      "GowidResponseListUserSearchResDto": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "type": "array",
            "items": {
              "$ref": "#/components/schemas/UserSearchResDto"
            }
          }
        }
      },
      "GowidResponsePurposeRequirementResDto": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/PurposeRequirementResDto"
          }
        }
      },
      "GowidResponsePurposeRequirementResDtoV2": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/PurposeRequirementResDtoV2"
          }
        }
      },
      "GowidResponseStatementBulkApproveResDto": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/StatementBulkApproveResDto"
          }
        }
      },
      "GowidResponseStatementBulkApproveResDtoV2": {
        "type": "object",
        "properties": {
          "result": {
            "$ref": "#/components/schemas/ResponseObject"
          },
          "totalCount": {
            "type": "integer",
            "format": "int64"
          },
          "data": {
            "$ref": "#/components/schemas/StatementBulkApproveResDtoV2"
          }
        }
      },
      "PurposeDto": {
        "type": "object",
        "properties": {
          "purposeId": {
            "type": "integer",
            "description": "용도 ID",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "용도 이름",
            "example": "식대"
          },
          "category": {
            "$ref": "#/components/schemas/CategoryVo"
          },
          "listOrder": {
            "type": "integer",
            "description": "용도 목록 정렬 순서 (오름차순)",
            "format": "int32"
          },
          "limitType": {
            "type": "string",
            "description": "한도 계산 방식 (PERSON=1인당, ITEM=결제 건당)",
            "enum": [
              "PERSON",
              "ITEM"
            ]
          },
          "limitAmount": {
            "type": "integer",
            "description": "한도 금액 (원화). 0이면 한도 미지정.",
            "format": "int64",
            "example": 100000
          },
          "isActivated": {
            "type": "boolean",
            "description": "용도 활성화 여부. false이면 사용 불가."
          },
          "hasRequirement": {
            "type": "boolean",
            "description": "세부입력항목(용도 필수 입력) 존재 여부"
          },
          "requirement": {
            "$ref": "#/components/schemas/PurposeRequirementDto"
          },
          "isDeducted": {
            "type": "boolean",
            "description": "용도에 설정된 공제 여부. true이면 해당 용도의 지출은 공제 처리됨."
          }
        }
      },
      "PurposeDtoV2": {
        "type": "object",
        "properties": {
          "purposeId": {
            "type": "integer",
            "description": "용도 id",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "용도 이름"
          },
          "category": {
            "$ref": "#/components/schemas/CategoryVo"
          },
          "listOrder": {
            "type": "integer",
            "description": "정렬 순서",
            "format": "int32"
          },
          "limitType": {
            "type": "string",
            "description": "한도 계산 방식",
            "enum": [
              "PERSON",
              "ITEM"
            ]
          },
          "limitAmount": {
            "type": "integer",
            "description": "한도 금액",
            "format": "int64"
          },
          "isActivated": {
            "type": "boolean",
            "description": "활성화 여부"
          },
          "requirements": {
            "type": "array",
            "description": "세부입력항목 목록",
            "items": {
              "$ref": "#/components/schemas/PurposeRequirementDtoV2"
            }
          },
          "deducted": {
            "type": "boolean"
          }
        }
      },
      "PurposeRequirementAnswerDtoV2": {
        "type": "object",
        "properties": {
          "purposeRequirementId": {
            "type": "integer",
            "description": "세부입력항목 id",
            "format": "int64"
          },
          "purposeRequirementName": {
            "type": "string",
            "description": "세부입력항목 이름"
          },
          "answers": {
            "type": "array",
            "description": "답변 목록",
            "items": {
              "type": "string",
              "description": "답변 목록"
            }
          }
        }
      },
      "PurposeRequirementDto": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer",
            "description": "세부입력항목 ID",
            "format": "int64"
          },
          "purposeId": {
            "type": "integer",
            "format": "int64",
            "writeOnly": true
          },
          "type": {
            "type": "string",
            "description": "응답 유형 (TEXT=텍스트, SELECT=단일선택, SELECT_MULTI=복수선택)",
            "enum": [
              "TEXT",
              "SELECT",
              "SELECT_MULTI"
            ]
          },
          "item": {
            "type": "string",
            "description": "세부입력항목 이름",
            "example": "방문 목적"
          },
          "options": {
            "uniqueItems": true,
            "type": "array",
            "description": "선택 옵션 목록. V1에서는 응답에 포함되지 않음.",
            "writeOnly": true,
            "items": {
              "type": "string",
              "description": "선택 옵션 목록. V1에서는 응답에 포함되지 않음."
            }
          },
          "guideDesc": {
            "type": "string",
            "description": "입력 안내 문구",
            "example": "거래처명을 입력하세요"
          },
          "isAvailableInput": {
            "type": "boolean",
            "description": "선택형 항목에서 직접 입력 허용 여부. type=SELECT 또는 SELECT_MULTI일 때 유효."
          }
        },
        "description": "세부입력항목 정보. hasRequirement=true일 때만 존재."
      },
      "PurposeRequirementDtoV2": {
        "type": "object",
        "properties": {
          "id": {
            "type": "integer",
            "description": "세부입력항목 id",
            "format": "int64"
          },
          "type": {
            "type": "string",
            "description": "응답 종류",
            "enum": [
              "TEXT",
              "SELECT",
              "SELECT_MULTI"
            ]
          },
          "item": {
            "type": "string",
            "description": "이름"
          },
          "guideDesc": {
            "type": "string",
            "description": "안내 문구"
          },
          "isAvailableInput": {
            "type": "boolean",
            "description": "종류가 선택인 경우, 직접입력 가능 여부"
          },
          "isRequired": {
            "type": "boolean",
            "description": "필수여부"
          }
        }
      },
      "PurposeRequirementResDto": {
        "type": "object",
        "properties": {
          "content": {
            "type": "array",
            "description": "세부입력항목 선택 옵션 목록",
            "items": {
              "type": "string",
              "description": "세부입력항목 선택 옵션 목록"
            }
          }
        }
      },
      "PurposeRequirementResDtoV2": {
        "type": "object",
        "properties": {
          "content": {
            "type": "array",
            "description": "세부입력항목 내용 목록",
            "items": {
              "type": "string",
              "description": "세부입력항목 내용 목록"
            }
          }
        }
      },
      "PurposeSimpleDto": {
        "type": "object",
        "properties": {
          "purposeId": {
            "type": "integer",
            "description": "용도 ID",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "용도 이름"
          },
          "limitType": {
            "type": "string",
            "description": "한도 계산 방식 (PERSON=1인당, ITEM=결제 건당)",
            "enum": [
              "PERSON",
              "ITEM"
            ]
          },
          "limitAmount": {
            "type": "integer",
            "description": "한도 금액 (원화). 0이면 한도 미지정.",
            "format": "int64",
            "example": 100000
          },
          "isActivated": {
            "type": "boolean",
            "description": "용도 활성화 여부. false이면 사용 불가."
          },
          "hasRequirement": {
            "type": "boolean",
            "description": "세부입력항목(용도 필수 입력) 존재 여부"
          }
        },
        "description": "지출 용도"
      },
      "PurposeSimpleDtoV2": {
        "type": "object",
        "properties": {
          "purposeId": {
            "type": "integer",
            "description": "용도 id",
            "format": "int64"
          },
          "name": {
            "type": "string",
            "description": "용도 이름"
          },
          "limitType": {
            "type": "string",
            "description": "한도 계산 방식",
            "enum": [
              "PERSON",
              "ITEM"
            ]
          },
          "limitAmount": {
            "type": "integer",
            "description": "한도 금액",
            "format": "int64"
          },
          "requirements": {
            "type": "array",
            "description": "세부입력항목 목록",
            "items": {
              "$ref": "#/components/schemas/PurposeRequirementDtoV2"
            }
          }
        },
        "description": "지출 사용 용도"
      },
      "PurposeVo": {
        "type": "object",
        "properties": {
          "name": {
            "type": "string",
            "description": "용도 이름"
          },
          "category": {
            "$ref": "#/components/schemas/CategoryVo"
          },
          "limitAmount": {
            "type": "integer",
            "description": "한도 금액",
            "format": "int64"
          },
          "listOrder": {
            "type": "integer",
            "description": "정렬 순서",
            "format": "int32"
          },
          "isActivated": {
            "type": "boolean",
            "description": "활성화 여부"
          },
          "hasRequirement": {
            "type": "boolean",
            "description": "세부 입력 항목 존재 여부"
          },
          "limitType": {
            "type": "string",
            "description": "한도 계산 방식",
            "enum": [
              "PERSON",
              "ITEM"
            ]
          }
        },
        "description": "지출 용도"
      },
      "ResponseObject": {
        "type": "object",
        "properties": {
          "code": {
            "type": "integer",
            "format": "int32"
          },
          "desc": {
            "type": "string"
          }
        }
      },
      "RoleSimpleResDto": {
        "type": "object",
        "properties": {
          "type": {
            "type": "string",
            "description": "권한 코드 (ROLE_MASTER=슈퍼 관리자, ROLE_MANAGER=지출 총괄 관리자, ROLE_VIEWER=세무 대리인, ROLE_MEMBER=일반 사용자, ROLE_CUSTOM=카드 관리자)",
            "enum": [
              "ROLE_MASTER",
              "ROLE_MANAGER",
              "ROLE_VIEWER",
              "ROLE_MEMBER",
              "ROLE_CUSTOM"
            ]
          },
          "name": {
            "type": "string",
            "description": "권한 이름 (예: 슈퍼 관리자)",
            "example": "슈퍼 관리자"
          },
          "description": {
            "type": "string",
            "description": "권한 설명"
          }
        },
        "description": "사용자 권한 정보"
      },
      "StatementBulkApproveResDto": {
        "type": "object",
        "properties": {
          "succeedStatements": {
            "type": "array",
            "description": "승인 성공한 지출 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseSimpleDto"
            }
          },
          "failedStatements": {
            "type": "array",
            "description": "승인 실패한 지출 목록",
            "deprecated": true,
            "items": {
              "$ref": "#/components/schemas/ExpenseSimpleDto"
            }
          }
        }
      },
      "StatementBulkApproveResDtoV2": {
        "type": "object",
        "properties": {
          "succeedStatements": {
            "type": "array",
            "description": "승인 성공한 지출 목록",
            "items": {
              "$ref": "#/components/schemas/ExpenseSimpleDtoV2"
            }
          },
          "failedStatements": {
            "type": "array",
            "description": "승인 실패한 지출 목록",
            "deprecated": true,
            "items": {
              "$ref": "#/components/schemas/ExpenseSimpleDtoV2"
            }
          }
        }
      },
      "UserSearchResDto": {
        "type": "object",
        "properties": {
          "userId": {
            "type": "integer",
            "description": "사용자 ID",
            "format": "int64"
          },
          "userName": {
            "type": "string",
            "description": "사용자 이름"
          },
          "email": {
            "type": "string",
            "description": "이메일 주소"
          },
          "isContractor": {
            "type": "boolean",
            "description": "법인 최초 계약자 여부"
          },
          "status": {
            "type": "string",
            "description": "사용자 상태 (PENDING=신청중, NORMAL=재직, PASSWORD_LOCK=잠김, ADD_ONLY=초대전, INVITATION=초대, INACTIVE=비활성, DELETED=삭제)",
            "enum": [
              "PENDING",
              "NORMAL",
              "PASSWORD_LOCK",
              "ADD_ONLY",
              "INVITATION",
              "INACTIVE",
              "DELETED"
            ]
          },
          "department": {
            "$ref": "#/components/schemas/DepartmentResDto"
          },
          "position": {
            "type": "string",
            "description": "직급",
            "example": "과장"
          },
          "role": {
            "$ref": "#/components/schemas/RoleSimpleResDto"
          },
          "notificationOnOff": {
            "type": "boolean",
            "description": "앱 알림 수신 여부"
          }
        }
      },
      "UserVo": {
        "type": "object",
        "properties": {
          "userName": {
            "type": "string",
            "description": "이름"
          },
          "email": {
            "type": "string",
            "description": "이메일"
          },
          "mobileNumber": {
            "type": "string",
            "description": "휴대폰번호"
          },
          "isInvitedUser": {
            "type": "boolean",
            "description": "초대여부"
          },
          "isContractor": {
            "type": "boolean",
            "description": "최초계약자 여부"
          },
          "isActivated": {
            "type": "boolean",
            "description": "활성사용자 여부"
          },
          "position": {
            "type": "string",
            "description": "직급"
          },
          "activatedAt": {
            "type": "string",
            "description": "활성일 (yyyyMMddHHmmss)"
          },
          "deactivatedAt": {
            "type": "string",
            "description": "비활성일 (yyyyMMddHHmmss)"
          },
          "notificationOnOff": {
            "type": "boolean",
            "description": "알림설정 여부"
          },
          "status": {
            "type": "string",
            "description": "사용자 상태",
            "enum": [
              "PENDING",
              "NORMAL",
              "PASSWORD_LOCK",
              "ADD_ONLY",
              "INVITATION",
              "INACTIVE",
              "DELETED"
            ]
          }
        }
      }
    },
    "securitySchemes": {
      "ApiKeyAuth": {
        "type": "apiKey",
        "description": "발급받은 API Key를 `Authorization` 헤더에 그대로 넣습니다. `Bearer` 같은 접두사는 붙이지 않습니다.",
        "name": "Authorization",
        "in": "header"
      }
    }
  }
}
```
