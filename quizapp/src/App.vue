<template>
  <div class="container">
    <Loading
      v-if="isLoading"
      :questions="questions"
      :numberOfQuestions="numberOfQuestions"
    ></Loading>
    <Score
      :score="score"
      :numberOfQuestions="numberOfQuestions"
      v-if="isActive"
    ></Score>
    <Picture
      v-if="isActive"
      :imgSrc="questions[currentQuestionIndex].photo_url"
    ></Picture>
    <Selector
      v-if="isActive"
      @submitAnswer="checkAnswer"
      @nextQuestion="nextQuestion"
      :questionAnswered="questionAnswered"
      :isAnswerCorrect="isAnswerCorrect"
      :question="questions[currentQuestionIndex]"
    ></Selector>
    <Override
      v-if="isActive"
      @overrideScore="overrideScore"
      :question="questions[currentQuestionIndex]"
      :questionAnswered="questionAnswered"
      :isAnswerCorrect="isAnswerCorrect"
    ></Override>
    <QuestionIndicator
      v-if="isActive"
      :currentQuestionIndex="currentQuestionIndex"
      :numberOfQuestions="numberOfQuestions"
    ></QuestionIndicator>
    <Finished v-if="isFinished" :score="score"></Finished>
  </div>
</template>

<script>
import Picture from "./components/Picture.vue";
import axios from "axios";
import Loading from "./components/Loading.vue";
import Selector from "./components/Selector.vue";
import Finished from "./components/Finished.vue";
import Score from "./components/Score.vue";
import Override from "./components/Override.vue";
import QuestionIndicator from "./components/QuestionIndicator.vue";

export default {
  name: "App",
  components: {
    Picture,
    Loading,
    Selector,
    Finished,
    Score,
    Override,
    QuestionIndicator,
  },
  data() {
    return {
      questions: [], // array to hold questions returned from the API
      questionCount: 0, // count the number of questions returned while fetching - used for loading page
      numberOfQuestions: 5, // number of questions to fetch
      currentQuestionIndex: 0, // index of current question
      score: 0,
      isAnswerCorrect: true,
      questionAnswered: false,
    };
  },
  // fetch questions on page load
  mounted() {
    for (
      this.questionCount = 0;
      this.questionCount < this.numberOfQuestions;
      this.questionCount++
    ) {
      axios
        .get("http://127.0.0.1:4000/getQuestions/test/1")
        .then((response) => this.questions.push(response.data.questions[0]));
    }
  },
  methods: {
    checkAnswer(answerManufacturer, answerModel) {
      // assign current question to new local variable
      var currentQuestion = this.questions[this.currentQuestionIndex];

      // check manufacturer
      if (currentQuestion.manufacturer === answerManufacturer) {
        this.score += 20;
      } else {
        this.isAnswerCorrect = false;
      }

      // check model
      if (currentQuestion.model === answerModel) {
        console.log("model correct " + answerModel);
        this.score += 80;
      } else {
        this.isAnswerCorrect = false;
      }
      this.questionAnswered = true;
    },
    // progress to next question
    nextQuestion() {
      if (this.currentQuestionIndex != this.numberOfQuestions) {
        this.isAnswerCorrect = true;
        this.currentQuestionIndex++;
        this.questionAnswered = false;
      }
    },
    overrideScore() {
      this.score += 100;
    },
  },
  computed: {
    isLoading() {
      return this.questions.length < this.numberOfQuestions;
    },
    isActive() {
      return (
        this.questions.length == this.numberOfQuestions &&
        this.currentQuestionIndex != this.numberOfQuestions
      );
    },
    isFinished() {
      return this.currentQuestionIndex === this.numberOfQuestions;
    },
  },
};
</script>

<style>
/* #app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
} */
</style>
