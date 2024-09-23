import { createStore } from 'vuex'

const store = createStore({
  state: { // 프로젝트에서 공통으로 사용할 변수 정의
    timerCategories: [],  // 카테고리 저장
    recordedTimeBlocks: [] // 기록된 타임 블록
  },
  // state를 변경하는 역할, 
  // 동기 처리를 함
  // (위의 함수가 실행되고 종료된 후 그 다음 아래의 함수가 실행됨)
  mutations: { 
    // 카테고리 추가
    addCategory(state, category) {
      state.timerCategories.push(category);
    },
    // 타임 블록 기록
    recordTimeBlock(state, { categoryIndex, timeBlockIndex }) {
        state.recordedTimeBlocks[timeBlockIndex] = state.timerCategories[categoryIndex].color;
    },
    // 타이머 업데이트 (경과 시간 저장)
    updateTimer(state, { index, elapsedTime }) {
      state.timerCategories[index].timer += elapsedTime;
    },
    // 카테고리 상태 초기화
    resetTimer(state, index) {
      state.timerCategories[index].timer = 0;
      state.timerCategories[index].timerRunning = false;
    }
  },
  actions: { // mutation을 실행시키는 역할, 비동기 처리를 함
    addCategory({ commit }, category) {
      commit('addCategory', category);
    },
    // 타임 블록 기록 액션
    recordTimeBlock({ commit }, { categoryIndex, timeBlockIndex }) {
        commit('recordTimeBlock', { categoryIndex, timeBlockIndex });
    },
    updateTimer({ commit }, { index, elapsedTime }) {
      commit('updateTimer', { index, elapsedTime });
    },
    resetTimer({ commit }, index) {
      commit('resetTimer', index);
    }
  },
  getters: { // 각 components의 계산된 속성의 공통 사용 정의
    getCategories: state => state.timerCategories,
    getTimeBlocks: state => state.recordedTimeBlocks
  }
});

export default store;