
<style scoped>

.game-container {
  max-width: 100%;
  margin: 0 auto;
  padding: 1rem;
  text-align: center;
}

h1 {
  color: #42b883;
  margin-bottom: 1rem;
}


.maze-grid {
  display: inline-block;
  margin: 1rem auto;
  border: 2px solid #2c3e50;
  background: white;
}

.maze-row {
  display: flex;
}

.maze-cell {
  width: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  border: 1px solid #eee;
  position: relative;
}


.wall {
  background-image: url('@/assets/pared.png');
  background-size: contain;
  background-position: center;
  background-repeat: no-repeat;
  color: transparent;
  font-size: 0; 
}

.path {
  background-color: #f8f9fa;
}

.start {
  background-color: #42b883;
  color: white;
}

.end {
  background-image: url('@/assets/meta.png');
  background-size: contain;
  background-position: center;
  background-repeat: no-repeat;
  color: transparent;
  font-size: 0;
}

.player {
  background-image: url('@/assets/character.png');
  background-size: contain;
  background-position: center;
  background-repeat: no-repeat;
  width: 100%;
  height: 100%;
  position: absolute;
}

.visited {
  background-color: rgba(255, 136, 0, 0.866); 
}

.solution-path {
  background-color: rgba(52, 152, 219, 0.5);
}

.best-path {
  background-color: #42b883;
}

.worst-path {
  background-color: #e74c3c;
}


.controls-group {
  display: flex;
  flex-direction: column;
  gap: 2rem; 
  margin-top: 2rem;
}

.controls {
  display: flex;
  justify-content: center;
  flex-wrap: wrap;
  gap: 1rem;
}

.path-navigation {
  margin: 1.5rem 0;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  padding: 0.8rem;
  background-color: #2c3e50;
  border-radius: 8px;
}

.path-counter {
  font-weight: bold;
  min-width: 150px;
}


button {
  padding: 0.8rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}


.nav-button {
  padding: 0.6rem 1rem;
  background-color: #3498db;
  color: white;
}

.start-button {
  background-color: #3498db;
  color: white;
}

.delete-button {
  background-color: #3498db;
  color: white;
}

.solution-button {
  background-color: #3498db;
  color: white;
}

.back-button {
  background-color: #3498db;
  color: white;
}

.save-button {
  background-color: #3498db;
  color: white;
}

.best-path-button {
  padding: 0.6rem 1rem;
  background-color: #3498db;
  color: white;
}

.path-animation-button {
  padding: 0.6rem 1rem;
  background-color: #3498db;
  color: white;
}

.worst-path-button {
  padding: 0.6rem 1rem;
  background-color: #3498db;
  color: white;
}


button:hover,
.nav-button:hover:not(:disabled),
.best-path-button:hover,
.worst-path-button:hover,
.start-button:hover,
.delete-button:hover,
.solution-button:hover,
.back-button:hover,
.save-button:hover {
  opacity: 0.9;
  transform: translateY(-2px);
}

.nav-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}


.message {
  margin-top: 1rem;
  padding: 0.8rem;
  border-radius: 6px;
  display: inline-block;
}

.error-message {
  background-color: #e74c3c;
  color: white;
}

.save-message {
  background-color: #00ff44;
  color: white;
}

</style>