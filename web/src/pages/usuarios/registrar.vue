<template>
  <div class="q-pa-md">
    <Atras />

    <q-form @submit="onSubmit" class="">
      <div class="row justify-center q-mt-md">
        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-input v-model.trim="input.nombres" outlined label="Nombres:" lazy-rules dense
            :rules="[(v) => (v && v.length > 0) || 'Dato requerido.', (v) => (v && v.length < 30) || '30 max permitido.',]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-input v-model.trim="input.apellidos" outlined label="Apellidos:" lazy-rules dense
            :rules="[(v) => (v && v.length > 0) || 'Dato requerido.', (v) => (v && v.length < 30) || '30 max permitido.',]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-input v-model.trim="input.documento" outlined label="N Documento:" lazy-rules dense :rules="[
            (v) => ((id && id.length != 0) || (v && v.length > 0)) || 'Dato requerido.',
            (v) => ((id && id.length != 0) || (v && v.length < 30)) || '30 max permitido.'
          ]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <vue3-q-tel-input v-model:tel="input.telefono" dense outlined required default-country="bo"
            :eager-validate="false" search-text="buscar pais..." no-results-text="no existe." :rules="[
              (v: any) => (v && v.length > 0) || 'Dato requerido.',
              (v: any) => (v && v.length < 20) || '20 max permitido.',
            ]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-input v-model.trim="input.username" outlined label="Username:" lazy-rules dense
            :rules="[(v) => (v && v.length > 0) || 'Dato requerido.', (v) => (v && v.length < 25) || '25 max permitido.',]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-input v-model.trim="input.password" class="q-mb-sm" outlined label="Password:" lazy-rules dense
            :hint="id ? 'Si deja vacio, se mantiene el valor anterior' : ''" :rules="[
              (v) => ((id && id.length != 0) || (v && v.length > 0)) || 'Dato requerido.',
              (v) => ((id && id.length != 0) || (v && v.length < 100)) || '100 max permitido.'
            ]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-input v-model.trim="input.domicilio" outlined label="Domicilio:" lazy-rules dense :rules="[
            (v) => !v || v.length === 0 || v.length >= 3 || '3 caracteres como minimo.',
            (v) => !v || v.length === 0 || v.length < 100 || '100 max permitido.'
          ]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-input v-model.trim="input.correo" outlined label="Correo:" lazy-rules dense clearable :rules="[
            (v) => !v || v.length === 0 || v.length >= 3 || '3 caracteres como minimo.',
            (v) => !v || v.length === 0 || v.length < 100 || '100 max permitido.',
            (v) => !v || v.length === 0 || v.length >= 3 && /.+@.+\..+/.test(v) || 'Correo electrónico inválido.'
          ]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-input v-model.trim="input.fecha_nac" outlined label="Fecha nacimiento:" type="date" lazy-rules dense
            clearable :rules="[
              (v) => !v || v.length === 0 || v.length >= 3 || '3 caracteres como minimo.',
              (v) => !v || v.length === 0 || v.length < 20 || '20 max permitido.'
            ]" />
        </div>

        <div class="col-xs-12 col-sm-6 q-px-md">
          <q-table title="Roles del usuario:" :loading="loading" :rows="roles" :columns="columns" row-key="name"
            hide-pagination :rows-per-page-options="[0]" :filter="filter" class="q-mx-md">

            <template v-slot:body-cell-check="props">
              <q-td :props="props">
                <q-checkbox dense v-model="props.row.check" label="" color="green" />
              </q-td>
            </template>
          </q-table>
        </div>
      </div>



      <div class="row justify-center">
        <q-btn outline color="green" :loading="store.is_loading_page" :label="id ? 'Guardar Cambios' : 'Registrar'"
          size="small" icon="check" class="q-mt-xl" type="submit" />
      </div>
    </q-form>


  </div>
</template>

<!-- eslint-disable @typescript-eslint/no-explicit-any -->
<script lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { pageStore } from '../../stores/pageStore';
import { RolesResponseCustom } from '../../types/roles';
import { useRoute, useRouter } from 'vue-router';
import { InstanciaNew, NewUsuario, UpdateUsuario, resetInstanciaNew } from '../../types/usuarios';
import 'vue3-q-tel-input/dist/vue3-q-tel-input.esm.css';
import Vue3QTelInput from 'vue3-q-tel-input';
import Atras from '../utils/header-back.vue'
import RolesService from '../../services/rolesService';
import UsuariosService from '../../services/usuariosService';

const columns = [
  { name: 'id', label: 'ID', field: 'id', sortable: true },
  { name: 'nombre', label: 'nombre', field: 'nombre', sortable: true },
  { name: 'check', label: 'Seleccione', field: 'check', sortable: true },
];

export default {
  name: 'IndexRoles',
  components: { Atras, Vue3QTelInput },

  watch: {
    'input.nombres': function () { this.generateUser() },
    'input.apellidos': function () { this.generateUser() },
    'input.documento': function () { this.generatePass() }
  },

  setup() {
    const store = pageStore();
    const loading = ref(false);
    const route = useRoute();
    const id = ref()
    const router = useRouter();
    const input = ref<NewUsuario>(InstanciaNew);
    const roles = ref<RolesResponseCustom[]>([]);
    const roleservice = new RolesService();
    const usuariosService = new UsuariosService();

    const generateUser = () => {
      if (id.value) return;
      const ini = input.value.nombres[0];
      const ape = input.value.apellidos.toLowerCase().replace(/\s+/g, '');
      let username = ini + ape;
      if (!ini) username = ape;
      let sinAcentos = username.normalize('NFD').replace(/[\u0300-\u036f]/g, '');
      let sinEnie = sinAcentos.replace(/ñ/g, 'n').replace(/Ñ/g, 'N');
      let resultado = sinEnie.toLowerCase();
      input.value.username = resultado;
    }
    const generatePass = () => {
      if (id.value) return;
      let doc = input.value.documento;
      doc = doc.toLowerCase();
      input.value.password = doc;
    }

    async function onSubmit() {
      input.value.roles = [];
      roles.value.forEach(x => {
        if (x.check) input.value.roles.push(x.id)
      });
      if (input.value.roles.length == 0) {
        store.notificacion('Seleccione al menos 1 rol', '', 'negative');
        return;
      }
      // console.log(input.value);
      const xinput: NewUsuario = JSON.parse(JSON.stringify(input.value));
      if (xinput.fecha_nac) {
        xinput.fecha_nac = xinput.fecha_nac + 'T00:00:00-04:00'
      }

      if (id.value) {
        const upd: UpdateUsuario = {
          id: id.value,
          nombres: xinput.nombres,
          apellidos: xinput.apellidos,
          username: xinput.username,
          password: xinput.password,
          telefono: xinput.telefono,
          foto_url: xinput.foto_url,
          correo: xinput.correo,
          documento: xinput.documento,
          domicilio: xinput.domicilio,
          fecha_nac: xinput.fecha_nac,
          roles: xinput.roles,
        }
        const res = await usuariosService.updateUsuario(upd).then((e) => e).catch((e) => e)
        console.log(res);
        if (res.updateUsuario) router.back()
      } else {

        const res = await usuariosService.createUsuario(xinput).then((e) => e).catch((e) => e)
        console.log(res);
        if (res.createUsuario) router.back()
      }

    }

    onBeforeUnmount(() => resetInstanciaNew())

    onMounted(async () => {
      loading.value = true;
      id.value = route.query.iduser;
      const res = await roleservice.roles(false).then((e) => e).catch((e) => e);
      roles.value = res.roles ? res.roles : [];
      roles.value.forEach(x => x.check = false);

      if (id.value) {
        const uname = '' + route.query.uname;
        const res = await usuariosService.usuarioByUsername(uname).then((e) => e).catch((e) => e);
        input.value.nombres = res.usuarioByUsername.nombres;
        input.value.apellidos = res.usuarioByUsername.apellidos;
        input.value.username = res.usuarioByUsername.username;
        input.value.password = res.usuarioByUsername.password;
        input.value.correo = res.usuarioByUsername.correo;
        input.value.telefono = res.usuarioByUsername.telefono;
        input.value.documento = res.usuarioByUsername.documento;
        input.value.domicilio = res.usuarioByUsername.domicilio;
        input.value.fecha_nac = res.usuarioByUsername.fecha_nac;
        res.usuarioByUsername.roles.forEach((r: any) => {
          const exis = roles.value.find(e => e.id == r.id)
          if (exis) exis.check = true;
        })
      }

      loading.value = false

    });
    return {
      store,
      loading,
      filter: ref(''),
      id,
      input,
      columns,
      roles,
      generateUser,
      generatePass,
      onSubmit,
    };
  },
};
</script>
