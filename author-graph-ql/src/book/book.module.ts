import { BookEntity } from './entity/book.entity';
import { TypeOrmModule } from '@nestjs/typeorm';
import { BookService } from './book.service';
import { BookResolver } from './book.resolver';
import { Module } from '@nestjs/common';

@Module({
    imports: [TypeOrmModule.forFeature([BookEntity])],

    providers: [BookResolver, BookService]
})
export class BookModule {}
