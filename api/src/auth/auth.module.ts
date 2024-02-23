import { Module } from '@nestjs/common';
import { PassportModule } from '@nestjs/passport';
import { JWTStrategy } from './jwt.strategy';
import { Auth0Service } from './auth0.service';

@Module({
  imports: [PassportModule.register({ defaultStrategy: 'jwt' })],
  providers: [JWTStrategy, Auth0Service],
  exports: [PassportModule, JWTStrategy, Auth0Service],
})
export class AuthModule {}
