<script setup>
import { useRouter } from 'vue-router'
import { ref, onMounted, onUnmounted } from 'vue'
import { AlgoritmoNReinas } from '../../wailsjs/go/main/App'
import { EventsOn, EventsOff } from '../../wailsjs/runtime/runtime.js'

const router = useRouter()

const savedData = JSON.parse(sessionStorage.getItem('gameData') || '{}')
const board = ref(savedData?.board || []) 
const queens = ref(savedData?.queens)
const gameLog = ref([])
const isRunning = ref(false) 

const goHome = () => {
  sessionStorage.removeItem('gameData')
  router.push({ name: 'Home' })
}

const startAlgorithm = () => {
  if (isRunning.value) return; 
  isRunning.value = true;
  gameLog.value = [] 
  gameLog.value.push('Iniciando algoritmo...')
  AlgoritmoNReinas(board.value, queens.value)
}


onMounted(() => {
  EventsOn('board-update', (newBoardState) => {
    board.value = newBoardState
  })

  EventsOn('algorithm-log', (message) => {
    gameLog.value.push(message)
    const logContent = document.querySelector('.log-content');
    if (logContent) {
        logContent.scrollTop = logContent.scrollHeight;
    }
  })

  EventsOn('algorithm-finished', (wasSuccessful) => {
      isRunning.value = false
      gameLog.value.push(wasSuccessful ? '¡Algoritmo completado!' : 'No se encontró solución.')
  })
})

onUnmounted(() => {
    EventsOff('board-update')
    EventsOff('algorithm-log')
    EventsOff('algorithm-finished')
})

</script>

<template>
  <div class="page-container">
    <div class="game-panel">
      <div class="board-grid">
        <div v-for="(row, Row) in board" :key="Row" class="board-row">
          <div 
            v-for="(cell, Col) in row" 
            :key="Col" 
            class="board-cell" 
            :class="{
              'light-square': (Row + Col) % 2 === 0, 
              'dark-square': (Row + Col) % 2 !== 0,
              'blocked-cell': cell === 2
            }"
          >
            <span v-if="cell === 1" class="queen-symbol">♛</span>
          </div>
        </div>
      </div>
      <div class="controls">
        <button @click="startAlgorithm" class="start-algorithm-button" :disabled="isRunning">
          {{ isRunning ? 'Ejecutando...' : 'Empezar Algoritmo' }}
        </button>
        <button @click="goHome" class="back-button" :disabled="isRunning">
          Terminar
        </button>
      </div>
    </div>
    <div class="log-panel">
      <div class="log-header">
        Registro del Algoritmo
      </div>
      <div class="log-content">
        <div v-if="gameLog.length === 0" class="log-placeholder">
          Presiona "Empezar Algoritmo" para ver los pasos...
        </div>
        <div v-for="(message, index) in gameLog" :key="index" class="log-message">
          {{ message }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>

.page-container {
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  justify-content: center;
  min-height: 90vh;
  padding: 2rem;
  gap: 2rem;
}
.game-panel {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1.5rem;
}
.controls {
  display: flex;
  gap: 1rem;    
}
.controls button {
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
  font-weight: bold;
}
.controls button:hover:not(:disabled) {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}
.controls button:disabled {
    cursor: not-allowed;
    opacity: 0.6;
}
.start-algorithm-button {
  background-color: #42b883 ; 
  color: white;
}
.start-algorithm-button:hover:not(:disabled) {
  background-color: #3aa876;
}
.back-button {
  background-color: #f44336;
  color: white;
}
.back-button:hover:not(:disabled) {
  background-color: #d32f2f;
}
.log-panel {
  width: 300px; 
  border-radius: 10px;
  background-color: #fff;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.15);
  display: flex;
  flex-direction: column;
}
.log-header {
  background-color: #41d92a; 
  color: white;
  padding: 12px 15px;
  font-weight: bold;
  font-size: 1.1rem;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
  display: flex;
  align-items: center;
  gap: 10px;
}
.log-content {
  height: 400px;
  padding: 15px;
  overflow-y: auto;
  font-family: 'Courier New', Courier, monospace;
  font-size: 0.9rem;
}
.log-placeholder {
  color: #888;
  text-align: center;
  margin-top: 20px;
}
.log-message {
  padding: 6px 0;
  border-bottom: 1px solid #f0f0f0;
  color: #333;
}
.log-message:last-child { border-bottom: none; }
.board-grid {
  display: inline-block;
  border: 4px solid #4a3a2a;
  box-shadow: 0 10px 20px rgba(0, 0, 0, 0.2);
  line-height: 0;
}
.board-row { display: flex; }
.board-cell {
  width: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  transition: background-color 0.2s ease-in-out;
}
.light-square { background-color: #f0d9b5; }
.dark-square { background-color: #b58863; }
.queen-symbol {
  font-size: 40px;
  color: #2c3e50;
  text-shadow: 0 0 5px rgba(255, 255, 255, 0.5);
}

.blocked-cell {
  background-color: rgba(255, 0, 0, 0.4);
  box-shadow: inset 0 0 0 2px red;
}
</style>