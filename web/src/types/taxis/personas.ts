export interface UsuariosResponse {
  id?: string;
  nombres: string;
  apellidos: string;
  username: string;
  foto_url: string;
  telefono: string;
  documento: string;
  fecha_nac: string;
  domicilio: string;
  correo: string;
  registrado: string;
  estado: boolean;
  dataname: string;
  roles:Roles[]
}

interface Roles {
  id?: string;
  nombre: string;
}
