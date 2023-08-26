<template>
  <q-card v-if="objeto" class="" bordered flat>
    <q-card-section class="q-pb-none">
      <div class="text-h6">Otra información:</div>
    </q-card-section>
    <q-separator inset />
    <q-card-section class="q-pt-none">
      <div v-for="(valor, clave) in clon" :key="clave" class="q-ma-xs row">
        <template v-if="getclaveVal(clave)">
          <div class="col-4">
            <b> {{ parseClave(clave) }}: </b>
          </div>
          <div class="col-6">
            <span>
              {{ gettext(clave, valor) }}
            </span>
          </div>
        </template>
      </div>
    </q-card-section>
  </q-card>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { pageStore } from '../../stores/pageStore';

export default defineComponent({
  name: 'ViewObject',
  props: {
    objeto: { type: Object, default: null },
  },
  watch: {
    objeto: function () {
      this.clon = { ...this.objeto };
    },
  },

  setup() {
    const store = pageStore();
    const clon = ref({});

    function getclaveVal(clave: string) {
      const existe =
        clave != '__typename' &&
        clave != 'fondo_certificado_base64' &&
        clave != 'locaciones' &&
        clave != 'can_add_users' &&
        clave != 'portada_base64';
      return existe;
    }

    function parseClave(c: string) {
      if (c == 'fecha_update') c = 'última_modificación';
      const cap = c.charAt(0).toUpperCase() + c.slice(1).toLowerCase();
      let fs = cap.split('_');
      let j = fs.join(' ');
      return j;
    }

    const gettext = (clave: string, valor: string) => {
      const b =
        clave == 'registrado' ||
        clave == 'fecha_registro' ||
        clave == 'fecha_update' ||
        clave == 'fecha_evento' ||
        clave == 'disponibilidad_a' ||
        clave == 'disponibilidad_b';
      return b ? store.parseDate(valor) : valor;
    };

    return {
      store,
      clon,
      parseClave,
      gettext,
      getclaveVal,
    };
  },
});
</script>
