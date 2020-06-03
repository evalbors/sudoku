<template>
  <div id="app">
		<table>
			<tr v-bind:key="row" v-for="(row, indexR) in board">
				<td v-bind:key="number" v-for="(number, indexC) in row" >
					<input type="text" name="" :value="number" :id="'row' + indexR + 'column' + indexC">
				</td>
			</tr>
		</table>
		<button @click="getBoard">Start Game</button>
  </div>
</template>

<script>
import "./assets/css/main.css";

export default {
  name: "app",
  data() {
    return {
      board: []
    };
  },
  methods: {
    getBoard: function() {
      window.backend.board().then(result => {
        this.board = result;
      });
    }
  }
};
</script>
