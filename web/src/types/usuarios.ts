import { Roles } from './roles';

export interface UsuarioLoginResponse {
  id: string;
  nombres: string;
  apellidos: string;
  username: string;
  foto_url: string;
  telefono: string;
  correo: string;
  documento: string;
  fecha_nac: string;
  domicilio: string;
  registrado: string;
  estado: boolean;
  dataname: string;
  roles: [Roles];
}

export interface UsuarioResponse {
  id: string;
  nombres: string;
  apellidos: string;
  username: string;
  foto_url: string;
  telefono: string;
  correo: string;
  documento: string;
  fecha_nac: string;
  domicilio: string;
  registrado: string;
  estado: boolean;
}

export interface NewUsuario {
  nombres: string;
  apellidos: string;
  username: string;
  password: string;
  telefono: string;
  foto_url: string;
  correo?: string;
  documento: string;
  fecha_nac?: string;
  domicilio?: string;
  roles: string[];
}
export interface UpdateUsuario {
  id: string;
  nombres: string;
  apellidos: string;
  username: string;
  password: string;
  telefono: string;
  foto_url: string;
  correo?: string;
  documento: string;
  fecha_nac?: string;
  domicilio?: string;
  roles: string[];
}

export const InstanciaNew: NewUsuario = {
  nombres: '',
  apellidos: '',
  username: '',
  password: '',
  telefono: '',
  foto_url: '',
  correo: undefined,
  documento: '',
  fecha_nac: undefined,
  domicilio: undefined,
  roles: [],
};

const resetInstanciaNe = { ...InstanciaNew };
export const resetInstanciaNew = () => {
  Object.assign(InstanciaNew, resetInstanciaNe);
};
