new Vue({
    el: '#app',
    data: {
        key: true,
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