<template>
  <div class="q-pa-md">
    <q-btn-group square flat outline class="q-mb-xl">
      <q-btn outline color="green" label="Agregar" size="small" icon="add" to="/adm/viajes/registrar" />
    </q-btn-group>

    <div class="row q-gutter-md">
      <div class="col-xs-12 col-sm-3 col-md-2">
        <q-card >
          <q-img src="https://cdn.quasar.dev/img/parallax2.jpg">
            <div class="absolute-bottom text-h6">
              Viajes Cercanos
            </div>
          </q-img>

          <q-card-section>
            <q-btn class="full-width" color="green" square to="/adm/viajes/viajesCercanos"> IR </q-btn>
          </q-card-section>
        </q-card>
      </div>

      <div class="col-xs-12 col-sm-3 col-md-2">
        <q-card >
          <q-img src="https://cdn.quasar.dev/img/parallax2.jpg">
            <div class="absolute-bottom text-h6">
              Viajes Por Usuario
            </div>
          </q-img>

          <q-card-section>
            <q-btn class="full-width" color="green" square to="/adm/viajes/viajesByUsuario"> IR </q-btn>
          </q-card-section>
        </q-card>
      </div>

      <div class="col-xs-12 col-sm-3 col-md-2">
        <q-card >
          <q-img src="https://cdn.quasar.dev/img/parallax2.jpg">
            <div class="absolute-bottom text-h6">
              Viajes Por Pasajero
            </div>
          </q-img>

          <q-card-section>
            <q-btn class="full-width" color="green" square to="/adm/viajes/viajesByPasajero"> IR </q-btn>
          </q-card-section>
        </q-card>
      </div>

      <div class="col-xs-12 col-sm-3 col-md-2">
        <q-card >
          <q-img src="https://cdn.quasar.dev/img/parallax2.jpg">
            <div class="absolute-bottom text-h6">
              Viajes Por Conductor
            </div>
          </q-img>

          <q-card-section>
            <q-btn class="full-width" color="green" square to="/adm/viajes/viajesByConductor"> IR </q-btn>
          </q-card-section>
        </q-card>
      </div>

      <div class="col-xs-12 col-sm-3 col-md-2">
        <q-card >
          <q-img src="https://cdn.quasar.dev/img/parallax2.jpg">
            <div class="absolute-bottom text-h6">
              Viajes Por Categoria
            </div>
          </q-img>

          <q-card-section>
            <q-btn class="full-width" color="green" square to="/adm/viajes/viajesByCategoria"> IR </q-btn>
          </q-card-section>
        </q-card>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { pageStore } from '../../../stores/pageStore';

export default {
  name: 'IndexEmpresas',
  setup() {
    const store = pageStore();

    return {
      store,
    };
  },
};
</script>
