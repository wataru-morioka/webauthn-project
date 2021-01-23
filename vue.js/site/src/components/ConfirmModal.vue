<template lang='pug'>
  div#confirmation-modal
    div#confirm-modal(class="ui basic modal confirm")
      div(class="content")
        p {{ confirmMessage }}
        //- p(v-model='confirmMessage')
      div(class="actions")
        div(class="ui red basic cancel inverted button")
          i(class="remove icon") No
        div#ok-button(class="ui green ok inverted button")
          i(class="checkmark icon") Yes
</template>

<script lang='ts'>
import { Component, Vue, Prop } from 'vue-property-decorator';

@Component
export default class ConfirmModal extends Vue {
  @Prop() private confirmMessage!: string;
}
</script>

<style scoped lang='scss'>
#confirmation-modal {
  position: fixed;
}

.content {
  text-align: center;

  p {
    font-size: 14px;
  }
}
</style>
