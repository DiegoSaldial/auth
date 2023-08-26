
<template>
  <q-page class="items-center justify-evenly">
    <div class="row justify-center">
      <div class="col-xl-12 col-md-6 col-lg-4 q-mt-xl">
        <p class="text-center text-h5">Bienvenido al menu principal</p>
      </div>
    </div>

    <div class="row justify-center">
      <div class="col-xl-12 col-md-7 col-lg-4 col-xl-3">
        <p class="text-center q-pa-md">
          Damos un cálido y cordial recibimiento al corazón de nuestra plataforma: el Menú Principal. Aquí, te invitamos a
          embarcarte en un viaje culinario y de opciones, donde tus deseos y necesidades son la brújula que guía cada
          selección. Al cruzar el umbral del Menú Principal, te sumergirás en un mundo de posibilidades diseñadas para
          satisfacer tus gustos más exquisitos y tus preferencias más arraigadas
        </p>
      </div>
    </div>

    <div class="row justify-center">
      <div class="col-xl-12 col-md-7 col-lg-5 col-xl-3 q-mb-xl text-center">
        <p class="text-center q-pa-md">
          Nuestro Menú Principal no solo es un camino hacia la variedad; es un pasaporte hacia la comodidad. Aquí,
          encontrarás todas las herramientas que necesitas para navegar sin esfuerzo por las profundidades de nuestra
          plataforma. Ya sea que estés buscando respuestas, entretenimiento o inspiración, cada entrada en el Menú
          Principal te acerca un paso más a tu objetivo.
        </p>

        <q-btn outline color="green" label="Ver Landing Page" size="small" icon="arrow_left" @click="$router.push('/')"
          class="q-mb-xl" />
      </div>
    </div>
  </q-page>
</template>


<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'IndexPage',
  components: {},
  setup() {
    return {};
  }
});
</script>
