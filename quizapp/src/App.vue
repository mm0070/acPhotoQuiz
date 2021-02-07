<template>
  <div>
    <!-- <Picture v-if="questions.length != 0" :imgSrc="questions[0].photo_url"></Picture> -->
    <Loading :questions="questions"></Loading>
  </div>
</template>

<script>
// import Picture from './components/Picture.vue'
import axios from 'axios';
import Loading from './components/Loading.vue';

export default {
  name: 'App',
  components: {
    // Picture,
    Loading
  },
  data() {
    return {
      questions: [], // array to hold questions returned from the API
      questionCount: 0, // count the number of questions returned while fetching - used for loading page
      numberOfQuestions: 3, // number of questions to fetch
    }
  },
  methods: {
  },
  mounted() {
    for (this.questionCount = 0; this.questionCount < this.numberOfQuestions; this.questionCount++) {
      axios.get("http://127.0.0.1:4000/getQuestions/1").then(response => this.questions.push(response.data.questions[0]));
    }
    }
  }

</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
