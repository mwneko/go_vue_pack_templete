<template>
  <v-app>

    <v-main>
      <orderPage/>
    </v-main>
  </v-app>
</template>

<script>
import orderPage from './components/order';

export default {
  name: 'App',

  components: {
    orderPage,
  },

  data: () => ({
    //
  }),
};
</script>
