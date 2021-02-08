<template>
  <div class="fixed">
    <button
      v-if="overrideAvailable && !overrideTriggered"
      @click="showDetails"
      type="button"
      class="btn btn-warning"
    >
      Override
    </button>
    <button
      v-if="overrideAvailable && overrideTriggered"
      @click="override"
      type="button"
      class="btn btn-danger"
    >
      Override
    </button>
  </div>
</template>

<script>
export default {
  name: "Override",
  props: ["questionAnswered", "isAnswerCorrect", "question"],
  data() {
    return {
      overrideTriggered: false,
    };
  },
  computed: {
    overrideAvailable() {
      return this.questionAnswered && !this.isAnswerCorrect;
    },
  },
  methods: {
    override() {
      this.$emit("overrideScore");
      this.overrideTriggered = false;
    },
    showDetails() {
      console.log(this.question);
      this.overrideTriggered = true;
    },
  },
};
</script>

<style scoped>
div.fixed {
  position: fixed;
  bottom: 5px;
  right: 5px;
}
</style>
