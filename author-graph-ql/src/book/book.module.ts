import { BookService } from './book.service';
import { BookResolver } from './book.resolver';
import { Module } from '@nestjs/common';

@Module({
    providers: [BookResolver, BookService]
})
export class BookModule {}
