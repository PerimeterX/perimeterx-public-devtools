FROM perl:5
RUN cpanm --notest --skip-satisfied MaxMind::DB::Writer Text::CSV_XS
RUN mkdir /out /app
WORKDIR /app
COPY build.pl ./
VOLUME "/out"
ENTRYPOINT ["perl", "build.pl"]
