<template>
    <div>
        <p>Loading questions: {{ loadingProgress }}/3</p>
    </div>
</template>

<script>
export default {
    name: 'Loading',
    props: ['questions'],
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