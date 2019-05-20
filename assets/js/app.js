new Vue({
    el: '#app',
    data: {
        key: true,
        testScore: 50,
        deviationValue: 50.0,
        averageScore: 50.0,
    },
    methods: {
        onChange(event) {
            if (event.target.id === "ts") {
                this.key = true;
            } else {
                this.key = false;
            }
        }
    },
});