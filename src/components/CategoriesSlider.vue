<template>
  <div class="categories-slider">
    <div class="categories-slider__container">
      <h2 class="title">Explore by category</h2>
      <div class="categories-slider-block">
        <swiper
          :slidesPerView="8"
          :spaceBetween="30"

          :navigation="true"
          class="mySwiper w-1/4"
        >
          <swiper-slide
            class="categories-slider-slide"
            v-for="(item, index) in 10"
            :key="index"
          >{{ index }}
          </swiper-slide>
        </swiper>
      </div>
    </div>
  </div>
</template>

<script>
// Import Swiper Vue.js components
import { Swiper, SwiperSlide } from "swiper/vue";

// swiper bundle styles
import "swiper/swiper-bundle.min.css";

// swiper core styles
import "swiper/swiper.min.css";

// modules styles
import "swiper/components/navigation/navigation.min.css";

// import Swiper core and required modules
import SwiperCore, { Navigation } from "swiper";

// install Swiper modules
SwiperCore.use([ Navigation]);
export default {
  components: {
    Swiper,
    SwiperSlide
  }
};
</script>
<style scoped lang="sass">
@import "../assets/sass/helps/mixins"
.categories-slider-block
  margin-bottom: 120px
  margin-right: 0
  width: auto
  height: 315px
  display: flex
  flex-direction: column


.categories-slider
  &-slide
    height: 250px
    width: 250px
    background-color: #fff
    color: #000
    text-align: center
    display: flex
    flex-direction: column
    align-items: center
    justify-content: center

    &:hover
      background-color: #7d7d7d
      transform: scale(1.2)
      color: #fff
      height: 315px

.title
  @extend %heading-three
  margin-bottom: 60px
  margin-top: 70px


</style>
