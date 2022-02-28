<template>
  <div class="first-view">
    <my-header />
    <hero />
  </div>
  <catelogies-slider />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import MyHeader from "@/components/MyHeader.vue";
import Hero from "@/components/Hero.vue";
import CatelogiesSlider from "@/components/CatelogiesSlider.vue";

export default defineComponent({
  name: "home",
  components: {
    MyHeader,
    Hero,
    CatelogiesSlider
  }
});
</script>
