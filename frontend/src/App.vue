<template>
  <div id="app">
    <div class="game">
		<table>
			<tr v-bind:key="row" v-for="(row, indexR) in board">
				<td v-bind:key="number" v-for="(number, indexC) in row" >
					<input type="text" name="" :value="number" :id="'row' + indexR + 'column' + indexC">
				</td>
			</tr>
		</table>
		<button @click="getBoard">Start Game</button>
    </div>
    <div class="test">
      <input v-model="name" @keyup.enter="sayHello">
      <h1>{{test}}</h1>
    </div>
  </div>
</template>

<script>
import "./assets/css/main.css";

export default {
  name: "app",
  data() {
    return {
      board: [],
      test: "",
      name: ""
    };
  },
  methods: {
    getBoard: function() {
      window.backend.board().then(result => {
        this.board = result;
      });
    },
    sayHello: function() {
      window.backend.test(this.name).then(result => {
        this.test = result;
      });
    }
  }
}

</script>
