<template>
  <div id="home-page">
    <h2 id="today">{{ formattedDate }}<br>오늘의 Lifelog</h2>
    <img 
      id="log-icon" 
      :src="logIcon"
      @click="openModal" 
      alt="Log Icon"
    />

    <!-- 모달 컴포넌트 호출 -->
    <HomePageModal v-if="isModalOpen" @close="closeModal" />

    <!-- 7x24 블록을 테이블로 구현 -->
    <table id="home-page-block">
      <tr v-for="(row, rowIndex) in 24" :key="rowIndex">
        <td>{{ rowIndex + 1 }}</td>
        <td v-for="col in 6" :key="col" :style="{ backgroundColor: getBlockColor(rowIndex, col) }"></td>
      </tr>
    </table>
  </div>
</template>

<script>
import { mapGetters} from 'vuex';
import { format } from 'date-fns';
import HomePageModal from './HomePageModal.vue';

export default {
  name: 'HomePage',
  components: {
    HomePageModal
  },
  data() {
    return {
      formattedDate: '',
      isModalOpen: false,
      timeBlocks: [], // 시간을 기록할 배열
      logIcon: new URL('@/assets/Log_icon.png', import.meta.url).href  // require 대신 import.meta 사용
    };
  },
  computed: {
    ...mapGetters(['getTimeBlocks']),
  },
  created() {
    const today = new Date();
    this.formattedDate = format(today, 'yyyy.MM.dd');
  },
  methods: {
    openModal() {
      this.isModalOpen = true;
    },
    closeModal() {
      this.isModalOpen = false;
    },
    getBlockColor(row, col) {
      const timeBlockIndex = row * 6 + col; // 시간 블록 인덱스 계산
      return this.getTimeBlocks[timeBlockIndex] || '#FFFFFF'; // 기록된 색상 또는 기본 색상
    }
  }
};
</script>

<style scoped>
#home-page {
  width: 100%;
  height: 100%;    
}
#today {
  position: absolute;
  width: 197px;
  height: 85px;
  left: 16px;
  top: 50px;
}

#log-icon {
  padding: 10px 17px;
  gap: 10px;
  position: absolute;
  width: 130px;
  height: 60px;
  left: calc(50% + 32px);
  top: 642px;
  cursor: pointer;
  z-index: 1;
}

/* 7x24 블록을 테이블로 구현한 스타일 */
#home-page-block {
  position: absolute;
  background: #ffffff;
  border-collapse: collapse;
  top: 150px;
  left: 20px;
  width: 350px;
  height: 300px;
}

#home-page-block td {
  border: 1px solid #b9b9b9;
  width: 50px;
  text-align: center;
}
</style>