<template>
    <div class="modal-overlay" @click="closeModal">
      <div class="modal-content" @click.stop>
        <div class="modal-header">
          <h3>{{ formattedDate }}의 lifelog</h3>
          <button>
            <img :src="require('@/assets/CalendarModal_close_icon.png')" @click="closeModal" alt="Close"/>
          </button>
        </div>
  
        <div class="modal-grid">
          <div class="image-text-section">
            <img class="modal-image" alt="Image"/>
            <p class="modal-text">Your text goes here</p>
          </div>
          <div class="block-section">
            <table>
              <tbody>
                <tr v-for="row in 9" :key="'row-' + row">
                  <td v-for="col in 8" :key="'col-' + col" class="block"></td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'CalendarModal', // 컴포넌트 이름 정의
    props: {
      selectedDate: {
        type: Date, // 부모 컴포넌트로부터 전달받는 prop, Date 객체 타입
        default: null, // 기본값을 null로 설정
      },
    },
    computed: {
      // 선택된 날짜를 형식화하여 문자열로 반환하는 계산된 속성
      formattedDate() {
        if (this.selectedDate && this.selectedDate instanceof Date) {
          // 선택된 날짜가 존재하고 유효한 Date 객체인 경우
          return `${this.selectedDate.getMonth() + 1}월 ${this.selectedDate.getDate()}일`;
        }
        // 선택된 날짜가 없거나 유효하지 않은 경우 대체 문자열 반환
        return '날짜를 선택하세요';
      },
    },
    methods: {
      // 모달을 닫기 위해 부모 컴포넌트에 이벤트를 전달하는 메소드
      closeModal() {
        this.$emit('close'); // 'close' 이벤트를 부모 컴포넌트로 전달
      },
    },
  };
</script>

  
  <style scoped>
  button {
    background: transparent;
    border: none;
    margin: 0;
    padding: 0;
  }
  
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100vw;
    height: 100vh;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }
  
  .modal-content {
    position: absolute;
    width: 358px;
    height: 504px;
    background: #FFFFFF;
    border-radius: 16px;
    overflow: hidden; /* 콘텐츠가 넘치지 않도록 */
  }
  
  .modal-header {
    display: flex;
    justify-content: space-between;
    padding: 0px 10px;
    margin: 0 auto;
    border: none;
  }
  
  .modal-grid {
    padding: 0px 20px;
    margin: 0 auto;
    display: flex;
    flex-direction: column;
  }
  
  .image-text-section {
    display: flex;
    align-items: center;
  }
  
  .modal-image {
    width: 52px; /* 이미지 크기 조정 */
    height: 73px;
    object-fit: cover;
    border-radius: 4px; /* 이미지의 모서리를 둥글게 */
    margin-right: 10px; /* 이미지와 텍스트 사이의 간격 */
  }
  
  .modal-text {
    width: 257px; /* 텍스트 영역의 너비 */
    height: 73px; /* 텍스트 영역의 높이 */
    background-color: #f0f0f0;
    border-radius: 16px; /* 텍스트 영역의 모서리를 둥글게 */
    padding: 0px 10px; /* 텍스트 영역의 안쪽 여백 */
    flex: 1;
  }
  
  .block-section {
    width: 100%;
    height: 100%;
  }
  
  table {
    width: 317px;
    height: 324px;
    border-spacing: 0; /* 셀 사이의 간격 제거 */
    border-collapse: collapse; /* 테두리 겹침 제거 */
  }
  
  td {
    background: #f5f5f5; /* 블록의 배경 색상 */
    border: 1px solid #b9b9b9; /* 블록의 테두리 */
    width: 10%; /* 8열의 너비 설정 */
    height: 10%; /* 9행의 높이 설정 */
    box-sizing: border-box; /* 테두리와 패딩을 포함하여 사이즈 계산 */
  }
  </style>
  