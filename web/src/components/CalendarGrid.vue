<template>
    <div class="calendar-grid">
      <div class="day-name" v-for="day in dayNames" :key="day">{{ day }}</div>
      <div class="day" 
        v-for="day in daysInMonth" 
        :key="day.date" 
        @click="handleClick(day.date)" 
        :class="{ today: isToday(day.date) }">
        <span v-if="day.date">{{ day.date.getDate() }}</span>
        <img v-if="day.date" :src="getImageForDay(day.date)" alt="이미지" class="day-image"/>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'CalendarGrid', // 컴포넌트 이름
    props: {
      // 부모 컴포넌트에서 전달받는 prop들
      dayNames: Array, // 요일 이름을 담은 배열 (일, 월, 화 등)
      daysInMonth: Array // 해당 월의 날짜 배열을 전달받음
    },
    methods: {
      // 특정 날짜를 클릭했을 때 실행되는 함수
      handleClick(date) {
        // 부모 컴포넌트에 'open-modal' 이벤트를 전달하고, 클릭된 날짜를 함께 전달함
        this.$emit('open-modal', date);
      },
      // 날짜에 해당하는 이미지 URL을 생성하는 함수
      getImageForDay(date) {
        // 예시로 날짜 숫자를 포함한 간단한 이미지 URL을 반환
        return `https://via.placeholder.com/50?text=${date.getDate()}`;
      },
      // 전달받은 날짜가 오늘인지 확인하는 함수
      isToday(date) {
        const today = new Date(); // 현재 날짜
        // 날짜 객체의 연도, 월, 일과 오늘의 연도, 월, 일이 같은지 비교하여 오늘인지 확인
        return (
          date &&
          date.getDate() === today.getDate() &&
          date.getMonth() === today.getMonth() &&
          date.getFullYear() === today.getFullYear()
        );
      }
    }
  };
</script>

  
  
  <style scoped>
  .calendar-grid {
    display: grid;
    grid-template-columns: repeat(7, 1fr);
    gap: 10px;
  }
  
  .day-name {
    text-align: center;
    font-weight: bold;
  }
  
  .day {
    text-align: center;
    border: 1px solid #FFFFFF;
    min-height: 80px;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    align-items: center;
    overflow: hidden;
  }
  
  .day.today {
    background-color: #e0f7fa;
  }
  
  .day-image {
    width: 40px;
    height: 74px;
    object-fit: cover;
    border-radius: 7px;
  }
  </style>
  