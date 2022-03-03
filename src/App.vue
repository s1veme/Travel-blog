<template>
  <router-view />
</template>
<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  name: "App"
});
</script>
<style lang="sass">
.first-view
  min-height: 100vh
  background: center/ cover no-repeat url('./assets/bg-hero-dark.jpg')

.categories-slider-block .swiper-wrapper
  align-items: center
  height: 315px
</style>
