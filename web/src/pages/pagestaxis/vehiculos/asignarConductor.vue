<template>
  <div class="q-pa-md">
    <Atras />

    <q-form @submit="onSubmit" class="q-mt-lg">
      <div class="row q-gutter-sm justify-center">

        <div class="col-xs-12 col-sm-6">
          <q-select v-model="conductor" :options="usuarios" option-label="nombres" option-value="id" outlined label="Seleccione conductor:" lazy-rules dense :rules="[
            (v) => (v && v.id > 0) || 'Dato requerido.',]">
            <template v-slot:option="scope">
              <q-item v-bind="scope.itemProps">
                <q-item-section avatar>
                  <q-icon name="bookmark" color="green" />
                </q-item-section>
                <q-item-section>
                  <q-item-label>{{ scope.opt.id + ' - ' + scope.opt.nombres + ' ' + scope.opt.apellidos }} [{{ scope.opt.username }}] </q-item-label>
                  <q-item-label caption v-for="(r,key) in scope.opt.roles" :key="key">+ {{ r.nombre }}</q-item-label>
                </q-item-section>
              </q-item>
            </template>
          </q-select>
        </div>
      </div>

      <div class="row justify-center">
        <q-btn outline color="green" :loading="store.is_loading_page" :label="id ? 'Guardar Cambios' : 'Registrar'"
          size="small" icon="check" class="q-mt-xl" type="submit" />
      </div>
    </q-form>


  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import Atras from '../../utils/header-back.vue';
import VehiculosServices from '../../../services/taxis/vehiculos';
import UsuariosService from '../../../services/usuariosService'
import { pageStore } from '../../../stores/pageStore';
import { CreateConductorVehiculos } from '../../../types/taxis/vehiculos';

export default {
  name: 'IndexRoles',
  components: { Atras },

  setup() {
    const store = pageStore();
    const loading = ref(false);
    const route = useRoute();
    const router = useRouter();
    const vehiculosServices = new VehiculosServices();
    const usuariosService = new UsuariosService();
    const usuarios = ref([]);
    const conductor = ref();
    const id = ref()

    function onSubmit() {
      if(conductor.value) {
        const data:CreateConductorVehiculos = {
          usuario_id: ''+conductor.value.id,
          vehiculo_id: id.value
        }
        registrar(data);
      }
    }

    async function registrar(input: CreateConductorVehiculos) {
      const res = await vehiculosServices.asignarVehiculo(input).then((e) => e).catch((e) => e)
      console.log(res);
      if (res.asignarVehiculo) router.back();
    }

    onMounted(async () => {
      loading.value = true;
      id.value = route.query.id;
      const resr = await usuariosService.usuarios().then((e) => e).catch((e) => e);
      usuarios.value = resr.usuarios;
      loading.value = false

    });
    return {
      store,
      loading,
      filter: ref(''),
      id,
      conductor,
      usuarios,
      onSubmit,
    };
  },
};
</script>
