<template>
<div>
    <v-container>
        <v-spacer></v-spacer>
        <v-card>
            <v-card-title>
                Total amount:{{amount}}
                <v-spacer></v-spacer>

                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
            </v-card-title>
            <v-data-table :headers="headers" :items="order_list" :search="search" :items-per-page="5"></v-data-table>
        </v-card>
    </v-container>
</div>
</template>

<script>
/* eslint-disable*/
import axios from 'axios';

export default {
    name: 'HelloWorld',
    data() {
        return {
            amount: null,
            order_list: [],
            search: '',
            headers: [{
                    text: 'Order name',
                    align: 'start',
                    sortable: false,
                    value: 'order',
                },
                {
                    text: 'Customer Company',
                    value: 'company'
                },
                {
                    text: 'Customer name',
                    value: 'customer'
                },
                {
                    text: 'Order date',
                    value: 'date'
                },
                {
                    text: 'Delivered Amount',
                    value: 'amount'
                },
                {
                  text:'Total Amount',
                  value:'total'
                }

            ],
        }
    },
    methods:{
         getOrder(){
         axios.get('http://localhost:8080/sampleData').then(response => {
                this.album = response.data
            }).catch(function (error) {
                console.log(error);
            });

         },
         sum(){
 
         }
    },
      mounted () {
    axios
      .get('http://localhost:8080/sampleData')
      .then(response => (this.order_list = response.data))
  }
}
</script>
