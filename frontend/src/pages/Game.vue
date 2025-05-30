<script setup>
import { useRouter } from 'vue-router'
import { ref} from 'vue'
const router = useRouter()
const savedData = JSON.parse(sessionStorage.getItem('gameData'))
const board = ref(savedData?.board)
const goHome = () => {
  router.push({ name: 'Home' })
}
</script>

<template>
  <div class="game-container">
    <div class="board-grid">
      <div v-for="(row, Row) in board" :key="Row" class="board-row">
        <div v-for="(cell, Col) in row" :key="Col" class="board-cell" :class="{'empty': cell === 0}">
          <span v-if="cell === 1">♛</span>
        </div>
      </div>
    </div>
  </div>
  <div class="controls">
    <button @click="goHome" class="back-button">
      Terminar juego
    </button>
  </div>
</template>

<style scoped>
.game-container {
  max-width: 100%;
  margin: 0 auto;
  padding: 1rem;
  text-align: center;
}

.board-grid {
  display: inline-block;
  margin: 1rem auto;
  border: 2px solid #2c3e50;
  background: white;
}

.board-row {
  display: flex;
}

.board-cell {
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  border: 1px solid #eee;
  position: relative;
}

.empty {
  background-color: #f8f9fa;
}

.back-button {
  background-color: #3498db;
  color: white;
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  margin-top: 2rem;
}

.back-button:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

.controls {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 1rem;
}



button {
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

button:hover
.back-button:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

</style>