<template>
    <div class="calendar-header">
      <h2 @click="toggleEditMode" v-if="!isEditing">
        {{ year }}년 {{ monthNames[currentMonth] }}
      </h2>
  
      <div v-if="isEditing" class="edit-mode">
        <select v-model="selectedYear" @change="updateDate">
          <option v-for="year in years" :key="year" :value="year">
            {{ year }}년
          </option>
        </select>
  
        <select v-model="selectedMonth" @change="updateDate">
          <option v-for="(month, index) in monthNames" :key="index" :value="index">
            {{ month }}
          </option>
        </select>
  
        <button @click="toggleEditMode">확인</button>
      </div>
  
      <div>
        <button @click="$emit('prev-month')" class="arrow-button">&lt;</button>
        <button @click="$emit('next-month')" class="arrow-button">&gt;</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'CalendarHeader', // 컴포넌트 이름 정의
    props: {
      year: Number, // 부모로부터 받은 현재 연도
      currentMonth: Number, // 부모로부터 받은 현재 월 (0부터 11까지의 숫자)
      monthNames: Array, // 부모로부터 받은 월 이름 배열
      years: Array, // 선택할 수 있는 연도 리스트
      isEditing: Boolean // 현재 편집 모드 여부
    },
    data() {
      return {
        // 초기 값은 부모로부터 받은 현재 연도와 월을 사용
        selectedYear: this.year, 
        selectedMonth: this.currentMonth
      };
    },
    methods: {
      // 편집 모드를 토글하는 메소드
      toggleEditMode() {
        // 'toggle-edit-mode' 이벤트를 부모 컴포넌트로 전송하여 편집 모드 토글 요청
        this.$emit('toggle-edit-mode');
      },
      // 선택된 연도와 월을 업데이트하는 메소드
      updateDate() {
        // 'update-date' 이벤트를 발생시켜 선택된 연도와 월 정보를 부모로 전달
        this.$emit('update-date', { year: this.selectedYear, month: this.selectedMonth });
      }
    }
  };
</script>

  
  
  <style scoped>
  .calendar-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .arrow-button {
    background: none;
    border: none;
    font-size: 28px;
    cursor: pointer;
  }
  </style>
  