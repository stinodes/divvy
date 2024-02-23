import { Injectable } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import { GetUsersRequest, ManagementClient } from 'auth0';
import { CreateUserDto } from 'src/users/dto/create-user.dto';

@Injectable()
export class Auth0Service {
  private client: ManagementClient;

  constructor(private configService: ConfigService) {
    this.client = new ManagementClient({
      domain: (configService.get('AUTH0_URL') as string).replace(
        'https://',
        '',
      ),
      clientId: 'VKzR9MiyZIqtborwmwollcoLoCIjsZ9H',
      clientSecret:
        'cerHdzkoWn--47VLDB0EqyaoBZOs09wFqU1LoYpDS8wk5kGBDTzQ7RS5cxmYLo5t',
    });
  }

  findAllAuth0Users(params: GetUsersRequest) {
    return this.client.users.getAll(params).then((r) => r.data);
  }

  findOneAuth0User(id: string) {
    return this.client.users
      .get({
        id,
      })
      .then((r) => r.data);
  }

  createAuth0User(user: CreateUserDto) {
    const auth0User = {
      ...user,
      connection: 'Username-Password-Authentication',
    };
    return this.client.users.create(auth0User).then((r) => r.data);
  }
}
