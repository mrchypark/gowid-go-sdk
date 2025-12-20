curl --location 'https://openapi.gowid.com/v1/members'  --header 'Authorization: {API KEY}'

없음	서버 헬스 체크	GET
{{BASE_URL}}/actuator/health	path,response	
조회	법인에 소속된 사용자 정보 조회	GET
{{BASE_URL}}/v1/members	path,response	- 응답에서 일부 필드들이 제거 되고, 부서/권한 정보가 새롭게 정의되어 추가 됩니다.
조회	지출내역 단건 조회	GET
{{BASE_URL}}/v1/expenses/{expenseId}		- [2025-10-10] 이용내역 공제 관련 필드가 추가되었습니다.
조회	지출관리 내역에서 용도, 사용자, 메모에 대한 키워드 검색 및 조회	GET
{{BASE_URL}}/v1/expenses	path,request,response	- 요청 파라미터에서 memoName → memo 로 변경 됩니다.
- 요청 파라미터에서 month → startDate 로 변경 됩니다. (월단위 → 시작날짜단위)
- 응답에서 일부 필드가 제거 됩니다.
- 응답에서 용도의 ‘필수항목’ 정보가 추가 됩니다.
- 응답에서 용도의 ‘한도유형’ 정보가 추가 됩니다.
조회	미제출 영수증 개수 조회	N/A	미제공	- 미제출 영수증 목록조회 응답에 전체 갯수 필드가 추가 되어 갯수 조회 API 는 제거 됩니다.
조회	미제출 영수증 목록 조회	GET
{{BASE_URL}}/v1/expenses/not-submitted	path,response	- 용도의 ‘한도유형’ 정보가 응답에 추가 됩니다.
- 응답에서 추천 용도의 필수항목 정보가 추가 됩니다.
- 응답에서 전체 미제출 영수증 갯수, 전체 페이지 수, 마지막 페이지 여부가 추가 됩니다.
- 법인의 전체 미제출 영수증 목록이 조회 됩니다.
조회	법인에서 정한 사용용도 정책 목록 조회	GET
{{BASE_URL}}/v1/purposes	path,response	- 용도의 ‘활성여부’ 요청 파라미터가 추가 됩니다.
- 응답에서 용도의 ‘한도유형’ 정보가 추가 됩니다.
- 응답에서 용도의 ‘필수항목’ 정보가 추가 됩니다.
조회	법인에서 정한 사용용도의 선택형 필수항목값 목록조회	https://open-{{BASE_URL}}/v1/purposes/{purposeId}/requirements	NEW	- 새로 추가된 API 입니다.
- 용도의 선택형 필수항목값 목록조회로, 선택형 필수항목이 설정된 용도의 경우 설정된 항목값들을 조회할 수 있습니다.
- 영수증 제출 시 선택형 필수항목 용도의 필수항목값 선택 시 사용
수정	지출내역 수정 - 용도 정보	PUT
{{BASE_URL}}/v1/expenses/{expenseId}/purposes	path,request,response	- RequestBody 에 수정하고자 하는 용도가 필수항목 설정이 되어 있을 경우, 필수항목 정보가 추가 됩니다.
- 응답에서 일부 필드가 제거 됩니다.
- 응답에서 이용내역의 ‘필수항목’ 정보가 추가 됩니다.
수정	지출내역 수정 - 다건 용도 변경	PUT
{{BASE_URL}}/v1/expenses/purposes	path,request,response	- 용도의 필수항목이 설정되어 있는 경우 RequestBody 에 필수항목 정보가 추가 됩니다.
- 용도의 필수항목이 설정되어 있는 경우 응답에 필수항목 정보가 추가 됩니다.
수정	지출내역 수정 - 메모 정보	PUT
{{BASE_URL}}/v1/expenses/{expenseId}/memo	path	
수정	지출내역 수정 - 내/외부 참석자 정보	PUT
{{BASE_URL}}/v1/expenses/{expenseId}/participants	path	
수정	지출내역 수정 - 일반 정보	PUT
{{BASE_URL}}/v1/expenses/{expenseId}	path	- 용도의 필수항목이 설정되어 있는 경우 RequestBody 에 필수항목 정보가 추가 됩니다.
- 용도의 필수항목이 설정되어 있는 경우 응답에 필수항목 정보가 추가 됩니다.
수정	지출내역 수정 - 승인 상태 정보	PUT
{{BASE_URL}}/v1/expenses/{expenseId}/approval-status	path,request	- RequestBody 에서 approvedAt 필드가 추가 됩니다.
수정	지출내역 수정 - 다건 승인	PATCH
{{BASE_URL}}/v1/expenses/approval-status/approved	path,request,response	- RequestBody 에서 단건 승인 상태 변경의 형태를 array 로 전달 주시면 됩니다.
- 응답에서 성공한 이용내역과 실패한 이용내역을 전달 드립니다.
추가	지출내역 댓글 추가	POST
{{BASE_URL}}/v1/expenses/{expenseId}/comments	path,request	- RequestBody 내 expenseId 필드가 제거 되었습니다.
