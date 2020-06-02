<template>
    <div id="app">
        <table>
            <tr v-bind:key="row" v-for="(row, indexR) in board">
                <td v-bind:key="number" v-for="(number, indexC) in row" >
                    <input type="text" name="" :value="number" :id="'row' + indexR + 'column' + indexC">
                </td>
            </tr>
        </table>
        <button class="buttonStart">Start Game</button>
    </div>
</template>

<script>
import "./assets/css/main.css";

export default {
  name: "app",

  data () {
    return {
        board: [
            [1, 2, 3, 4, 5, 6, 7, 8, 9],
            [4, 5, 6, 7, 8, 9, 1, 2, 3],
            [7, 8, 9, 1, 2, 3, 4, 5, 6],
            [2, 1, 4, 3, 6, 5, 8, 9, 7],
            [3, 6, 5, 8, 9 ,7, 2, 1, 4],
            [8, 9, 7, 2, 1, 4, 3, 6, 5],
            [5, 3, 1, 6, 4, 2, 9, 7, 8],
            [6, 4, 2, 9, 7, 8, 5, 3, 1],
            [9, 7, 8, 5, 3, 1, 6, 4, 2]
        ]
    }
  },
};
</script>
