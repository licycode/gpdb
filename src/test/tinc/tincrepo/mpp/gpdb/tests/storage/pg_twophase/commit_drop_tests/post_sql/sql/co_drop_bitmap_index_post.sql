CREATE INDEX cr_co_bitmap_idx1 ON cr_co_table_bitmap_index USING bitmap (numeric_col);
insert into cr_co_table_bitmap_index select i||'_'||repeat('text',100),i,i||'_'||repeat('text',3),i,i,i,'{3}',i,i,i,'2006-10-19 10:23:54', '2006-10-19 10:23:54+02', '1-1-2002' from generate_series(101,200)i;
\d cr_co_table_bitmap_index
set enable_seqscan=off;
select numeric_col from cr_co_table_bitmap_index where numeric_col=1;
drop table cr_co_table_bitmap_index;
