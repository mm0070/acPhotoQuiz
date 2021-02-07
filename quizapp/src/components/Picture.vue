<template>
    <div>
        <img :src="imgSrc">
    </div>
</template>

<script>
export default {
    name: 'Picture',
    props: ['imgSrc']
}
</script>

<style scoped>

</style>