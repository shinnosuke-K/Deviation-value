Vue.component('demo-grid', {
    template: '#grid-template',
    delimiters: ['${', '}'],

    props: {
      heroes: Array,
      columns: Array,
      filterKey: String
    },
    data: function () {
      var sortOrders = {}
      this.columns.forEach(function (key) {
        sortOrders[key] = 1
      })
      return {
        sortKey: '',
        sortOrders: sortOrders
      }
    },
    computed: {
      filteredHeroes: function () {
        var sortKey = this.sortKey
        var filterKey = this.filterKey && this.filterKey.toLowerCase()
        var order = this.sortOrders[sortKey] || 1
        var heroes = this.heroes
        if (filterKey) {
          heroes = heroes.filter(function (row) {
            return Object.keys(row).some(function (key) {
              return String(row[key]).toLowerCase().indexOf(filterKey) > -1
            })
          })
        }
        if (sortKey) {
          heroes = heroes.slice().sort(function (a, b) {
            a = a[sortKey]
            b = b[sortKey]
            return (a === b ? 0 : a > b ? 1 : -1) * order
          })
        }
        return heroes
      }
    },
    filters: {
      capitalize: function (str) {
        return str.charAt(0).toUpperCase() + str.slice(1)
      }
    },
    methods: {
      sortBy: function (key) {
        this.sortKey = key
        this.sortOrders[key] = this.sortOrders[key] * -1
      }
    }
}),

new Vue({
    el: '#app',
    delimiters: ['${', '}'],
    data: {
        key: true,

        infoLists: [],

        testScore: 50,
        deviationValue: 50.0,
        averageScore: 50.0,

        inputScore: 50,
        inputValue: 50,

        resultDeviValue: 0,
        resultTestScore: 0,

        prefecture: '',
        prefectures: [
            { name: '北海道',   value: 'hokkaido'},
            { name: '青森県',   value: 'aomori'},
            { name: '岩手県',   value: 'iwate'},
            { name: '宮城県',   value: 'miyagi'},
            { name: '秋田県',   value: 'akita'},
            { name: '山形県',   value: 'yamagata'},
            { name: '福島県',   value: 'fukushima'},
            { name: '茨城県',   value: 'ibaraki'},
            { name: '栃木県',   value: 'tochigi'},
            { name: '群馬県',   value: 'gumma'},
            { name: '埼玉県',   value: 'saitama'},
            { name: '千葉県',   value: 'chiba'},
            { name: '東京都',   value: 'tokyo'},
            { name: '神奈川県', value: 'kanagawa'},
            { name: '新潟県',   value: 'niigata'},
            { name: '富山県',   value: 'toyama'},
            { name: '石川県',   value: 'ishikawa'},
            { name: '福井県',   value: 'fukui'},
            { name: '山梨県',   value: 'yamanashi'},
            { name: '長野県',   value: 'nagano'},
            { name: '岐阜県',   value: 'gifu'},
            { name: '静岡県',   value: 'shizuoka'},
            { name: '愛知県',   value: 'aiti'},
            { name: '三重県',   value: 'mie'},
            { name: '滋賀県',   value: 'siga'},
            { name: '京都府',   value: 'kyoto'},
            { name: '大阪府',   value: 'osaka'},
            { name: '兵庫県',   value: 'hyogo'},
            { name: '奈良県',   value: 'nara'},
            { name: '和歌山県', value: 'wakayama'},
            { name: '鳥取県',   value: 'tottori'},
            { name: '島根県',   value: 'shimane'},
            { name: '岡山県',   value: 'okayama'},
            { name: '広島県',   value: 'hiroshima'},
            { name: '山口県',   value: 'yamaguchi'},
            { name: '徳島県',   value: 'tokushima'},
            { name: '香川県',   value: 'kagawa'},
            { name: '愛媛県',   value: 'ehime'},
            { name: '高知県',   value: 'kochi'},
            { name: '福岡県',   value: 'hukuoka'},
            { name: '佐賀県',   value: 'saga'},
            { name: '長崎県',   value: 'nagasaki'},
            { name: '熊本県',   value: 'kumamoto'},
            { name: '大分県',   value: 'oita'},
            { name: '宮崎県',   value: 'miyazaki'},
            { name: '鹿児島県', value: 'kagoshima'},
            { name: '沖縄県',   value: 'okinawa'},
        ],

        searchQuery: '',
        gridColumns: ['name', 'power'],
        gridData: [
            { name: 'Chuck Norris', power: Infinity },
            { name: 'Bruce Lee', power: 9000 },
            { name: 'Jackie Chan', power: 7000 },
            { name: 'Jet Li', power: 8000 }
        ]
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
            var stanDeviation = 10 * (this.testScore - this.averageScore) / (this.deviationValue - 50)
            this.resultDeviValue = Math.ceil((10 * (this.inputScore - this.averageScore) / stanDeviation + 50) * 10) / 10
        },
        // calucate test Score
        onChangeValue() {
            var stanDeviation = 10 * (this.testScore - this.averageScore) / (this.deviationValue - 50)
            this.resultTestScore = Math.ceil(stanDeviation * (this.inputValue - 50) / 10 + this.averageScore)
        },

        onGetInfo() {
            if (this.key) {
                fetch('/info?prefecture=' + this.prefecture + "&deviation=" + Math.round(this.resultDeviValue))
                .then(response => response.json())
                .then(response => {
                    this.infoLists = response
                })
                .catch(error => console.error('Error:', error));
            } else {
                fetch('/info?prefecture=' + this.prefecture + "&deviation=" + Math.round(this.inputValue))
                .then(response => response.json())
                .then(response => {
                    this.infoLists = response
                })
                .catch(error => console.error('Error:', error));
            }
        }
    },
});