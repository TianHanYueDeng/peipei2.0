<template>
  <div class="informationFrame">
    <div class="personalData"></div>
    <div class="positionAndFood"></div>
  </div>
</template>
<script lang="ts">
import { Component, Vue, Prop } from "vue-property-decorator";
@Component({
  name: "Information"
})
export default class Information extends Vue {}
</script>
<style lang="stylus">
.informationFrame {
  display: flex;
  margin-top: 15vw;
  margin-bottom: 4vw;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  width: 84vw;
  height: 55.5vh;
}

.personalData {
  background-image: url('../../../public/information/personalData.png');
  background-size: 100% 100%;
  height: 40vw;
  width: 65vw;
}

.positionAndFood {
  background-image: url('../../../public/information/positionAndFood.png');
  background-size: 100% 100%;
  height: 76vw;
  width: 65vw;
}
</style>