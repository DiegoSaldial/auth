<template>
  <q-dialog v-model="alert" position="bottom" @hide="hide" style="width: 70vh;">
      <q-card>
        <q-card-section>
          <div class="text-h6">Marque la ubicacion y haga click en aceptar</div>
        </q-card-section>

        <q-card-section class="q-pt-none">
          <l-map
            v-if="showMap"
            :zoom="zoom"
            :center="center"
            :options="mapOptions"
            style="height: 50vh;width: 40vw;"
            @click="handleMapClick"
          >
            <l-tile-layer :url="url" :attribution="attribution"/>
            <l-marker :lat-lng="markerLatLng"></l-marker>
          </l-map>
        </q-card-section>

        <q-card-actions align="right">
          <q-btn flat label="Aceptar" color="green" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>
</template>

<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts">
import { defineComponent, ref } from 'vue';
import 'leaflet/dist/leaflet.css';
import { latLng } from 'leaflet';
import { LMap, LTileLayer,LMarker } from '@vue-leaflet/vue-leaflet';

export default defineComponent({
  name: 'IndexPage',
  components: {
    LMap,
    LMarker,
    LTileLayer,
  },
  setup(_,vue) {
    const refmap = ref();
    const alert = ref(false);
    const markerLatLng = ref<Float32Array[]>([]);

    function opendialog(lat:any,lon:any){
      alert.value = true;
      markerLatLng.value = [lat, lon];
    }

    function hide(){
      vue.emit('select',markerLatLng.value);
    }

    function handleMapClick(event:any) {
      const { lat, lng } = event.latlng;
      markerLatLng.value = [lat, lng];
    }

    return {
      alert,
      refmap,
      zoom: 14.2,
      center: latLng(-21.528098, -64.730105),
      url: 'https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png',
      attribution:'&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors',
      showMap: true,
      mapOptions: {
        zoomSnap: 0.5
      },
      markerLatLng,
      handleMapClick,
      hide,
      opendialog,
    };
  }
});
</script>
