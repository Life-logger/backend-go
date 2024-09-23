<template>
  <div class="calendar-container">
    <CalendarHeader
      :year="currentYear"
      :currentMonth="currentMonth"
      :monthNames="monthNames"
      :years="years"
      :isEditing="isEditing"
      @toggle-edit-mode="toggleEditMode"
      @prev-month="prevMonth"
      @next-month="nextMonth"
    />

    <CalendarGrid
      :dayNames="dayNames"
      :daysInMonth="daysInMonth"
      @open-modal="openModal"
    />

    <!-- CalendarModal 컴포넌트로 모달창 이동 -->
    <CalendarModal 
      v-if="showModal" 
      :selectedDate="selectedDate" 
      @close="closeModal" 
    />
  </div>
</template>

<script>
import CalendarHeader from './CalendarHeader.vue'; // 달력의 상단 헤더를 표시하는 컴포넌트
import CalendarGrid from './CalendarGrid.vue'; // 달력 날짜를 그리드 형태로 표시하는 컴포넌트
import CalendarModal from './CalendarModal.vue'; // 날짜 선택 후 나타나는 모달 창

export default {
  name: 'CalendarPage', // 컴포넌트 이름 정의
  components: {
    CalendarHeader, // 달력의 헤더 컴포넌트
    CalendarGrid, // 달력의 날짜를 표시하는 그리드 컴포넌트
    CalendarModal // 선택된 날짜의 상세 정보를 보여줄 모달 컴포넌트
  },
  data() {
    return {
      currentMonth: new Date().getMonth(), // 현재 월
      currentYear: new Date().getFullYear(), // 현재 연도
      isEditing: false, // 날짜 및 월을 편집할 수 있는 모드 여부
      selectedYear: new Date().getFullYear(), // 선택된 연도 (편집 모드용)
      selectedMonth: new Date().getMonth(), // 선택된 월 (편집 모드용)
      dayNames: ['일', '월', '화', '수', '목', '금', '토'], // 요일 이름
      monthNames: [
        '1월', '2월', '3월', '4월', '5월', '6월',
        '7월', '8월', '9월', '10월', '11월', '12월'
      ], // 월 이름 배열
      years: Array.from({ length: 101 }, (_, i) => new Date().getFullYear() - 50 + i), 
      // 현재 연도를 기준으로 101년 범위의 연도 배열 생성 (과거 50년부터 미래 50년까지)
      showModal: false, // 모달 창 표시 여부
      selectedDate: null, // 선택된 날짜 (모달에 표시할 날짜)
    };
  },
  computed: {
    // 현재 월의 일수를 계산하는 계산 속성
    daysInMonth() {
      const days = [];
      const firstDayOfMonth = new Date(this.currentYear, this.currentMonth, 1); // 해당 월의 첫 번째 날
      const lastDayOfMonth = new Date(this.currentYear, this.currentMonth + 1, 0); // 해당 월의 마지막 날

      const startDay = firstDayOfMonth.getDay(); // 해당 월의 첫날의 요일 인덱스
      // 첫 주의 시작 요일 전까지 빈 칸을 추가
      for (let i = 0; i < startDay; i++) {
        days.push({ date: null });
      }

      // 각 날짜를 추가
      for (let day = 1; day <= lastDayOfMonth.getDate(); day++) {
        days.push({
          date: new Date(this.currentYear, this.currentMonth, day), // 날짜 객체로 추가
        });
      }

      return days; // 그리드에 사용할 날짜 배열 반환
    },
    year() {
      return this.currentYear; // 현재 연도 반환
    },
  },
  methods: {
    // 이전 달로 이동하는 메소드
    prevMonth() {
      if (this.currentMonth === 0) { // 1월이면 이전 달은 12월, 연도를 1년 줄임
        this.currentMonth = 11;
        this.currentYear--;
      } else {
        this.currentMonth--; // 그 외에는 월을 1 감소
      }
    },
    // 다음 달로 이동하는 메소드
    nextMonth() {
      if (this.currentMonth === 11) { // 12월이면 다음 달은 1월, 연도를 1년 증가
        this.currentMonth = 0;
        this.currentYear++;
      } else {
        this.currentMonth++; // 그 외에는 월을 1 증가
      }
    },
    // 편집 모드를 토글하는 메소드
    toggleEditMode() {
      this.isEditing = !this.isEditing; // 편집 모드의 상태를 반전
      if (!this.isEditing) {
        // 편집 모드를 종료하면 선택된 연도와 월을 현재 연도와 월로 맞춤
        this.selectedYear = this.currentYear;
        this.selectedMonth = this.currentMonth;
      }
    },
    // 날짜를 업데이트하는 메소드
    updateDate({ year, month }) {
      this.currentYear = year; // 연도 업데이트
      this.currentMonth = month; // 월 업데이트
    },
    // 해당 날짜의 이미지를 가져오는 메소드 (샘플 이미지 URL)
    getImageForDay(date) {
      return `https://via.placeholder.com/50?text=${date.getDate()}`; // 날짜를 포함한 이미지 생성
    },
    // 모달을 여는 메소드
    openModal(date) {
      if (date) {
        this.selectedDate = date; // 선택된 날짜 저장
        this.showModal = true; // 모달 표시
      }
    },
    // 모달을 닫는 메소드
    closeModal() {
      this.showModal = false; // 모달 숨김
      this.selectedDate = null; // 선택된 날짜 초기화
    },
    // 해당 날짜가 오늘인지 확인하는 메소드
    isToday(date) {
      const today = new Date();
      return (
        date &&
        date.getDate() === today.getDate() && // 날짜 비교
        date.getMonth() === today.getMonth() && // 월 비교
        date.getFullYear() === today.getFullYear() // 연도 비교
      );
    },
  },
};
</script>

<style scoped>
  .calendar-container {
    width: 90vw; /* 화면 너비의 90%로 설정 */
    max-width: 358px; /* 최대 너비를 358px로 제한 */
    height: 90vh; /* 화면 높이의 90%로 설정 */
    max-height: 634px; /* 최대 높이를 634px로 제한 */
    margin: 0 auto; /* 화면 중앙에 배치 */
    background: #FFFFFF;
    border: 1px solid #FFFFFF;
    border-radius: 10px; /* 전체 달력 컨테이너의 모서리를 둥글게 */
    padding: 10px; /* 내부 여백 추가 */
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.1); /* 약간의 그림자 추가 */
    overflow: hidden; /* 내용이 넘칠 경우 숨김 */
  }
</style>
