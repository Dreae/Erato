package lib

import (
  "io"
  "os"
  "archive/tar"
  "archive/zip"
  "compress/gzip"
)

// Unzip recursively extracts files from provided Reader representing a
// ZIP archive
func Unzip(src io.Reader, size int64, dest string) error {
  r, err := zip.NewReader(src, size)
  if err != nil {
    return err
  }

  os.MkdirAll(dest, 0755)

  extractAndWriteFile := func(f *zip.File) error {
    rc, err := f.Open()
    if err != nil {
      return err
    }
    defer func() {
      if err := rc.Close(); err != nil {
        panic(err)
      }
    }()

    path := filepath.Join(dest, f.Name)

    if f.FileInfo().IsDir() {
      os.MkdirAll(path, f.Mode())
    } else {
      f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
      if err != nil {
        return err
      }
      defer f.Close()

      _, err = io.Copy(f, rc)
      if err != nil {
        return err
      }
    }
    return nil
  }

  for _, f := range r.File {
    err := extractAndWriteFile(f)
    if err != nil {
      return err
    }
  }

  return nil
}

// Untargz extracts files from a Reader representing a .tar.gz archive into
// dest directory
func Untargz(src io.Reader, dest string) error {
  gzr, err := gzip.NewReader(src)
  if err != nil {
    return err
  }

  r := tar.NewReader(gzr)
  for {
    header, err := r.Next()
    if err == io.EOF {
      break
    }

    if err != nil {
      return err
    }

    switch header.TypeFlag {
    case tar.TypeDir:
      os.MkdirAll(header.Name, header.Mode)
    case tar.TypeReg:
      f, err := os.OpenFile(header.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, header.Mode)
      if err != nil {
        return err
      }
      defer f.Close()

      _, err := io.Copy(f, r)
      if err != nil {
        return err
      }
    }
  }

  return nil
}
