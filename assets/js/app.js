new Vue({
    el: '#app',
    data: {
        key: true,

        testScore: 50,
        deviationValue: 50.0,
        averageScore: 50.0,

        inputScore: 50,
        inputValue: 50,

        resultDeviValue: 0,
        resultTestScore: 0,

    },
    methods: {
        onChange(event) {
            if (event.target.id === "ts") {
                this.key = true;
            } else {
                this.key = false;
            }
        },
        // calucate deviation value
        onChangeScore() {
            console.log(this.testScore, this.averageScore, this.deviationValue)
            var stanDeviation = 10 * (this.testScore - this.averageScore) / (this.deviationValue - 50)

            console.log(stanDeviation)
            this.resultDeviValue = Math.ceil((10 * (this.inputScore - this.averageScore) / stanDeviation + 50) * 10) / 10
            console.log(this.resultDeviValue)
        },
        // calucate test Score
        onChangeValue() {
            var stanDeviation = 10 * (this.testScore - this.averageScore) / (this.deviationValue - 50)
            this.resultTestScore = Math.ceil(stanDeviation * (this.inputValue - 50) / 10 + this.averageScore)
        }
    },
});