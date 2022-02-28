<template>
  <home />
</template>
<script lang="ts">
import { defineComponent } from "vue";
import Home from "@/views/Home.vue";

export default defineComponent({
  name: "App",
  components: { Home }

});
</script>
<style lang="sass">
.first-view
  min-height: 100vh
  background: center/ cover no-repeat url('./assets/bg-hero-dark.jpg')
</style>
