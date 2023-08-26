<template>
  <div class="q-pl-md q-pt-xs">
    <div class="row justify-start">
      <div class="col-12 col-md-2">
        <q-btn outline color="orange" label="Volver" size="small" icon="arrow_left" @click="$router.back()" />
      </div>
      <div class="col-12 col-md-9">
        <p class="text-center text-h5">{{ title }}</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'HeaderBack',
  props: {
    title: String,
  },
});
</script>
