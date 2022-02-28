<template>
  <div class="first-view">
    <my-header />
    <hero />
  </div>

  <categories-slider />
</template>

<script lang="ts">
import { defineComponent } from "vue";
import MyHeader from "@/components/MyHeader.vue";
import Hero from "@/components/Hero.vue";
import CategoriesSlider from "@/components/CategoriesSlider.vue";

export default defineComponent({
  name: "home",
  components: {
    MyHeader,
    Hero,
    CategoriesSlider
  }
});
</script>
