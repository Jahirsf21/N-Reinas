<script setup>
import { useRouter } from 'vue-router'
import { CreateBoard } from '../../wailsjs/go/main/App'
import { ref } from 'vue'

const router = useRouter()

const size = ref(null) 
const queens = ref(null)
const board = ref([])

const errorMessage = ref('')
/**
* Function that validates user input, generates the board from the backend
* and redirects to the game view if everything is correct.
*/
const startGame = async () => { 
  errorMessage.value = ''

  const boardSize = Number(size.value)
  const queenCount = Number(queens.value)

  if (isNaN(boardSize) || boardSize < 1 || boardSize > 25) {
    errorMessage.value = 'El tamaño del tablero debe ser un número entre 1 y 25.'
    return 
  }

  if (isNaN(queenCount) || queenCount <= 0) {
    errorMessage.value = 'La cantidad de reinas debe ser un número mayor a 0.'
    return
  }

  try {

    board.value = await CreateBoard(size.value)
    sessionStorage.setItem('gameData', JSON.stringify({
      board: board.value,
      queens: queens.value
    }))

    router.push({ name: 'Game' })

  } catch (error) {
    errorMessage.value = `Error al crear el tablero: ${error}`
    console.error(error)
  }
}
</script>

<template>
  <div class="principal-page">
    <h1>Bienvenido a N-Reinas</h1>
    <div v-if="errorMessage" class="error-message">
      {{ errorMessage }}
    </div>
    <div class="input-container">
      <input 
        v-model.number="size" 
        type="number" 
        placeholder="Dimensiones del tablero (1-25)" 
        class="board-input-size"
        min="1"
        max="25"
      >
      <input 
        v-model.number="queens" 
        type="number" 
        placeholder="Cantidad de reinas" 
        class="board-queens-number"
        min="1"
      >
    </div>
    <div class="buttons-container">
      <button @click="startGame" class="start-game">Empezar</button>
    </div>
  </div>
</template>

<style scoped>

.principal-page {
  text-align: center;
  padding: 2rem;
}

.buttons-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  margin-top: 20px;
  margin-bottom: 20px;
}

.start-game {
  padding: 1.5rem;
  background-color: #42b883;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 1.2rem;
  cursor: pointer;
  transition: all 0.3s ease;
  width: 250px;
  text-align: center;
}

.start-game:hover {
  background-color: #3aa876;
  transform: translateY(-3px);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.input-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
  margin-top: 2rem;
}

.board-input-size,
.board-queens-number {
  padding: 1rem;
  width: 250px;
  font-size: 1rem;
  border: 2px solid #ccc;
  border-radius: 8px;
  outline: none;
  transition: border-color 0.3s ease, box-shadow 0.3s ease;
  margin-top: 1rem;
}

.board-input-size:focus,
.board-queens-number:focus {
  border-color: #42b883;
  box-shadow: 0 0 8px rgba(66, 184, 131, 0.3);
}

.error-message {
  margin: 20px auto;
  padding: 10px;
  border-radius: 5px;
  background-color: #ff475720;
  color: #ff4757;
  font-weight: bold;
  max-width: 400px;
}
</style>