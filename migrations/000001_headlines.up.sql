BEGIN TRANSACTION ISOLATION LEVEL SERIALIZABLE;

CREATE TABLE data_sources(
    id SERIAL NOT NULL,

    title VARCHAR (255) NOT NULL,
    description VARCHAR (512) NOT NULL,
    link VARCHAR (512) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    UNIQUE(title),
    PRIMARY KEY (id)
);

CREATE TABLE headlines(
      id SERIAL NOT NULL,
      data_source_id INTEGER NOT NULL,

      title VARCHAR (255) NOT NULL,
      description VARCHAR (512) NOT NULL,
      link VARCHAR (512) NOT NULL,
      html_link VARCHAR (512) NOT NULL,
      author VARCHAR (255) NOT NULL,

      published_at TIMESTAMPTZ NOT NULL,

      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
      updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

      FOREIGN KEY (data_source_id) REFERENCES data_sources (id),

      UNIQUE(title),
      PRIMARY KEY (id)
);

END TRANSACTION;