<template>
    <div>
        <p>Loading questions: {{ loadingProgress }}/{{ numberOfQuestions }}</p>
    </div>
</template>

<script>
export default {
    name: 'Loading',
    props: ['questions', 'numberOfQuestions'],
    computed: {
        loadingProgress() {
            if (this.questions == undefined) {
                return "0";
            } else {
                return this.questions.length;
            }
        }
    }
}
</script>