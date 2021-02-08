<template>
    <div>
        <select class="form-select" v-model="manufacturerSelect">
            <option selected>Manufacturer</option>
            <option>Airbus</option>
            <option>Boeing</option>
            <option>Embraer</option>
            <option>Antonov</option>
            <option>ATR</option>
            <option>Bombardier</option>
            <option>Cessna</option>
            <option>Fokker</option>
            <option>Tupolev</option>
            <option>Ilyushin</option>
        </select>
        <select class="form-select" v-model="modelList">
            <option v-for="option in setOptions" :value="option.val" :key="option.val">{{ option.text }}</option>
        </select>
        <button type="button" class="btn btn-success" @click="submitAnswer">Next</button>
    </div>
</template>

<script>
export default {
    name: 'Selector',
    data() {
        return {
            manufacturerSelect: undefined,
            modelList: undefined,
        }
    },
    methods: {
        submitAnswer() {
            this.$emit('submitAnswer', this.manufacturerSelect, this.modelList)
        }
    },
    computed: {
        setOptions() {
            if (this.manufacturerSelect === 'Airbus') { 
                var modelList = [
                    {val: 'A300', text:'A300'},
                    {val: 'A310', text:'A310'},
                    {val: 'A318', text:'A318'},
                    {val: 'A319', text:'A319'},
                    {val: 'A320', text:'A320'},
                    {val: 'A321', text:'A321'},
                    {val: 'A330', text:'A330'},
                    {val: 'A340', text:'A340'},
                    {val: 'A350', text:'A350'},
                    {val: 'A380', text:'A380'},
                ]
            } else if (this.manufacturerSelect === 'Boeing') { 
                modelList = [
                    {val: '717', text:'717'},
                    {val: '727', text:'727'},
                    {val: '737-1', text:'737-100'},
                    {val: '737-2', text:'737-200'},
                    {val: '737-3', text:'737-300'},
                    {val: '737-4', text:'737-400'},
                    {val: '737-5', text:'737-500'},
                    {val: '737-6', text:'737-600'},
                    {val: '737-7', text:'737-700'},
                    {val: '737-8', text:'737-800'},
                    {val: '737-9', text:'737-900'},
                    {val: '8 MAX', text:'737-8 MAX'},
                    {val: '9 MAX', text:'737-9 MAX'},
                    {val: '10 MAX', text:'737-10 MAX'},
                    {val: '747-1', text:'747-100'},
                    {val: '747-2', text:'747-200'},
                    {val: '747SP', text:'747SP'},
                    {val: '747-3', text:'747-300'},
                    {val: '747-4', text:'747-400'},
                    {val: '747-8', text:'747-800'},
                    {val: '757-2', text:'757-200'},
                    {val: '757-3', text:'757-300'},
                    {val: '767-2', text:'767-200'},
                    {val: '767-3', text:'767-300'},
                    {val: '767-4', text:'767-400'},
                    {val: '777-2', text:'777-200'},
                    {val: '777-3', text:'777-300'},
                    {val: '787-8', text:'787-8'},
                    {val: '787-9', text:'787-9'},
                    {val: '787-10', text:'787-10'},
                ]
            } else if (this.manufacturerSelect === 'Embraer') { 
                modelList = [
                    {val:'E170', text:'E170'},
                    {val: 'E175', text:'E175'},
                    {val: 'E190', text:'E190'},
                    {val: 'E195', text:'E195'},
                    {val: 'ERJ-145', text:'ERJ-145'},
                    {val: 'EMB-120', text:'EMB-120'},
                ]
            } else if (this.manufacturerSelect === 'Antonov') { 
                modelList = [
                    {val:'AN-124', text:'AN-124'},
                    {val: 'AN-225', text:'AN-225'},
                ]
            } else if (this.manufacturerSelect === 'ATR') { 
                modelList = [
                    {val:'ATR-42', text:'ATR-42'},
                    {val: 'ATR-72', text:'ATR-72'},
                ]
            } else if (this.manufacturerSelect === 'Bombardier') { 
                modelList = [
                    {val:'Q400', text:'Q400/Dash-8'},
                    {val: 'CRJ-200', text:'CRJ-200'},
                    {val: 'CRJ-900', text:'CRJ-900'},
                    {val: 'CRJ-700', text:'CRJ-700'},
                    {val: 'Global', text:'Global'},
                ]
            } else if (this.manufacturerSelect === 'Cessna') { 
                modelList = [
                    {val:'152', text:'152'},
                    {val: '172', text:'172'},
                    {val: '182', text:'182'},
                    {val: 'Citation', text:'Citation'},
                ]
            } else if (this.manufacturerSelect === 'Fokker') { 
                modelList = [
                    {val:'50', text:'50'},
                    {val: '70', text:'70'},
                    {val: '100', text:'100'},
                ]
            } else if (this.manufacturerSelect === 'Tupolev') { 
                modelList = [
                    {val:'134', text:'Tu-134'},
                ]
            } else if (this.manufacturerSelect === 'Ilyushin') { 
                modelList = [
                    {val:'62' ,text:'Il-62'},
                ]
            }
            return modelList
        }
    }
}
</script>

<style scoped>

</style>