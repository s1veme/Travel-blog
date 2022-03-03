<template>
  <div class="authorization">
    <div class="authorization__container">
      <p class="info heading-two"> This page is intended for authorization of the administration. </p>
      <form class="signIn">
        <h2 class="authorization__title heading-three">Sign In</h2>
        <label for="signInEmail">Email</label>
        <input v-model="user.email" class="authorization__input" id="signInEmail" type="text" placeholder="Email...">
        <label for="signInPassword">Password</label>
        <input v-model="user.password" class="authorization__input" id="signInPassword" type="password"
               placeholder="Email...">
        <button
          @click="onSubmit"
          class="authorization__button button" type="button">Log in
        </button>
      </form>
      <form class="signUp">

      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import IUser from "@/types/IUser";

export default defineComponent({
  name: "Authorization",
  setup() {
    const user: IUser = reactive({
      email: "",
      password: ""
    });

    function onSubmit() {
      console.log("SEND JSON", JSON.stringify(user));

    };
    return {
      onSubmit,
      user
    };
  }
});
</script>

<style scoped lang="sass">
.authorization
  &__title
    margin-bottom: 1rem

  &__input
    border: none
    border-bottom: 2px solid #fff
    background-color: transparent
    color: #ffffff
    padding: 1rem 0px 1rem 0.5rem
    margin-bottom: 2rem

  &__button
    background-color: rgba(255, 255, 255, 0.80)
    margin: 0 auto
    padding: 1rem 2rem
    text-transform: uppercase
    border-radius: 5px
    font-size: 18px
    font-weight: bold
    transition: all 0.3s ease-in-out

    &:hover
      color: #ffffff
      background-color: rgba(255, 255, 255, 0.4)

.signIn
  position: fixed
  z-index: 999
  background-color: rgba(255, 255, 255, 0.1)
  padding: 30px
  left: 50%
  top: 50%
  transform: translate(-50%, -50%)
  width: 25%
  height: 40%
  display: flex
  flex-direction: column
  border-radius: 15px

.info
  text-align: center
  margin-top: 2rem
  color: rgba(255, 255, 255, 0.5)
  font-weight: 300
</style>