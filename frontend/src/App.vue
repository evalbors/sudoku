<template>
    <div id="app">
        <div class="start">
            <table>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row1" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row2" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row3" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row4" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row5" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row6" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row7" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row8" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
                <tr class="row">
                    <td v-bind:key="number" v-for="number in row" >
                        <input type="text" name="" :value="number">
                    </td>
                </tr>
            </table>
            <button class="buttonStart">Start Game</button>
            <!-- <div v-bind:key="number" v-for="number in numbers">
                {{number}}
            </div> -->
        </div>
    </div>
</template>

<script>
import "./assets/css/main.css";

export default {
  name: "app",

  data () {
    return {
        row1: [1, 2, 3, 4, 5, 6, 7, 8, 9],
        row2: [4, 5, 6, 7, 8, 9, 1, 2, 3],
        row3: [7, 8, 9, 1, 2, 3, 4, 5, 6],
        row4: [2, 1, 4, 3, 6, 5, 8, 9, 7],
        row5: [3, 6, 5, 8, 9 ,7, 2, 1, 4],
        row6: [8, 9, 7, 2, 1, 4, 3, 6, 5],
        row7: [5, 3, 1, 6, 4, 2, 9, 7, 8],
        row8: [6, 4, 2, 9, 7, 8, 5, 3, 1],
        row9: [9, 7, 8, 5, 3, 1, 6, 4, 2]
    }
  },
};
</script>
