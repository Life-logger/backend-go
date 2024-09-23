TimerModal.vue
<template>
  <div class="modal-overlay" @click="closeModal">
    <div class="modal-content" @click.stop>
      <h2>새 타이머 카테고리 추가</h2>
      <input v-model="categoryName" type="text" placeholder="카테고리 이름 입력" />

      <!-- 색상 선택 -->
      <input v-model="categoryColor" type="color" placeholder="카테고리 색상 선택" />

      <!-- 카테고리 저장 버튼 -->
      <button @click="saveCategory">저장</button>
      <button @click="closeModal">취소</button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'TimerModal',
  data() {
    return {
      categoryName: '',
      categoryColor: '#FFFFFF'
    };
  },
  methods: {
    saveCategory() {
      if (this.categoryName.trim() === '') {
        alert('카테고리 이름을 입력해주세요.');
        return;
      }

      const newCategory = {
        name: this.categoryName.trim(),
        color: this.categoryColor
      };

      this.$emit('save-category', newCategory); // 상위 컴포넌트로 카테고리 정보 전달
      this.closeModal();
    },
    closeModal() {
      this.$emit('close');
    }
  }
};
</script>

<style scoped>
.modal-overlay {
  position: absolute;
  width: 389px;
  height: 689px;
  left: 1px;
  top: 155px;
  background: #FFFFFF;
  border-radius: 16px 16px 0px 0px;
  z-index: 2; /* 높은 값으로 설정하여 우선순위 증가 */
  display: flex; /* 플렉스 박스로 설정 */
  justify-content: center; /* 가로 중앙 정렬 */
}

.modal-content {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  width: 300px;
  display: flex;
  flex-direction: column;
  align-items: center; /* 가로 중앙 정렬 */
}
.modal-content textarea,
.modal-content input[type="color"],
.modal-content button {
  margin-bottom: 15px;
  width: 100%;
}

.modal-content textarea {
  padding: 8px;
  height: 50px;
  resize: none;
  border-radius: 7px;
}

.modal-content input[type="color"] {
  height: 30px;
  border: none;
  padding: 0;
}
.modal-content button {
  background-size: contain; /* 이미지 크기를 버튼 크기에 맞춤 */
  background-repeat: no-repeat; /* 이미지 반복 안 함 */
  background-position: center; /* 버튼 중앙에 이미지 배치 */
  background: transparent; /* 배경색 투명으로 설정 */
  height: 56px;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>