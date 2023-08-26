/* eslint-disable @typescript-eslint/no-explicit-any */
import { defineStore } from 'pinia';
import { DocumentNode } from 'graphql/language/ast';
import { Opciones, InsOpciones, Site, InstanciaSite } from '../types/responses';
import { run_graphql_query } from './graphql/rungraphql-query';
import { run_graphql_mutation } from './graphql/rungraphql-mutation';
import { computed, ref } from 'vue';
import { mostrarTiempoRestante } from './graphql/sessionTime';
import { notificar } from './graphql/notificacion';

export const pageStore = defineStore('pageStore', () => {
  const sitio = ref<Site>(InstanciaSite);
  const loading_page = ref(false);

  const getsite = computed(() => {
    sitio.value.name = '' + process.env.VUE_APP_NAME;
    return sitio.value;
  });

  const is_loading_page = computed(() => loading_page.value);

  function setDataLogin(data: any) {
    data = JSON.stringify(data);
    window.localStorage.setItem('xdatalogin', data);
  }
  function getDataLogin() {
    const datalogi = window.localStorage.getItem('xdatalogin');
    let data = JSON.parse('' + datalogi);
    if (!data) data = {};
    return data;
  }

  async function setSessionTime(data: any, reload: boolean) {
    return await mostrarTiempoRestante(data, reload);
  }

  function logout() {
    window.localStorage.clear();
  }

  function parseDate(f: string, show_hours = true) {
    if (!f) return '-';
    if (!f.includes('T')) f += 'T::';
    const partes = f.split('T');
    const fs1 = partes[0].split('-');
    const fs2 = partes[1].split(':');
    if (show_hours) {
      return `${fs1[2]}/${fs1[1]}/${fs1[0]} ${fs2[0]}:${fs2[1]}`;
    } else {
      return `${fs1[2]}/${fs1[1]}/${fs1[0]}`;
    }
  }

  function formatCamelCaseString(input: string) {
    const words = input.split(/(?=[A-Z])/);
    const formattedWords = words.map(
      (word) => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase()
    );
    const formattedString = formattedWords.join(' ');
    return formattedString;
  }

  function notificacion(t: string, s = '', color = 'primary') {
    notificar(t, s, color);
  }

  const run_graphql = (
    sql: DocumentNode,
    variables = {},
    opciones: Opciones = InsOpciones
  ) => {
    return new Promise(async (resolve, reject) => {
      const mutar = sql.loc?.source.body.includes('mutation');
      if (mutar) {
        const res = await run_graphql_mutation(sql, variables, opciones)
          .then((e) => e)
          .catch((e) => e);
        // console.log('res', res);
        if (res.data) resolve(res.data);
        else reject(res);
      } else {
        const res = await run_graphql_query(sql, variables, opciones)
          .then((e) => e)
          .catch((e) => e);
        if (res.data) resolve(res.data);
        else reject(res);
      }
    });
  };
  return {
    getsite,
    is_loading_page,
    setSessionTime,
    run_graphql,
    logout,
    formatCamelCaseString,
    parseDate,
    notificacion,
    setDataLogin,
    getDataLogin,
  };
});
